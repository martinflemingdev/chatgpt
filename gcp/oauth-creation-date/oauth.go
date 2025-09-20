// go get google.golang.org/api/iam/v1 cloud.google.com/go/logging@v1
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/logging/logadmin"
	"google.golang.org/api/iam/v1"
)

func main() {
	ctx := context.Background()

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Fatal("set GOOGLE_CLOUD_PROJECT")
	}
	parent := fmt.Sprintf("projects/%s/locations/global", projectID)

	// IAM v1 (OAuth clients + secrets)
	iamSvc, err := iam.NewService(ctx)
	if err != nil {
		log.Fatalf("iam.NewService: %v", err)
	}
	oauthSvc := iam.NewProjectsLocationsOauthClientsService(iamSvc)
	credSvc := iam.NewProjectsLocationsOauthClientsCredentialsService(iamSvc)

	// Cloud Logging admin client (to read Admin Activity logs)
	logAdmin, err := logadmin.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("logadmin.NewClient: %v", err)
	}
	defer logAdmin.Close()

	fmt.Println("== OAuth 2.0 Clients (with creation times) ==")
	pageTok := ""
	for {
		resp, err := oauthSvc.List(parent).PageToken(pageTok).Do()
		if err != nil {
			log.Fatalf("oauth clients list: %v", err)
		}
		for _, c := range resp.OauthClients {
			clientCreate := findClientCreateTime(ctx, logAdmin, projectID, c)
			fmt.Printf("- %s\n  clientId=%s\n  displayName=%q  type=%s\n  created=%s\n",
				c.Name, c.ClientId, c.DisplayName, c.ClientType, tsOrNA(clientCreate))

			// list secrets for this client
			creds, err := credSvc.List(c.Name).Do()
			if err != nil {
				log.Printf("  (warn) list credentials for %s: %v", c.Name, err)
				continue
			}
			for _, cr := range creds.OauthClientCredentials {
				secretCreate := findSecretCreateTime(ctx, logAdmin, projectID, c, cr)
				fmt.Printf("    secret: %s  status=enabled:%t  created=%s\n",
					short(cr.Name), !cr.Disabled, tsOrNA(secretCreate))
			}
		}
		if resp.NextPageToken == "" {
			break
		}
		pageTok = resp.NextPageToken
	}
}

// findClientCreateTime returns the earliest Admin Activity timestamp for creating this OAuth client.
// It matches by client resource name and clientId in the log payloads.
func findClientCreateTime(ctx context.Context, admin *logadmin.Client, projectID string, c *iam.OauthClient) *time.Time {
	// We look for: iam.googleapis.com, method oauthClients.create
	// and try to narrow by both the client resource name and the clientId.
	filter := `
resource.type="audited_resource"
protoPayload.serviceName="iam.googleapis.com"
protoPayload.methodName="projects.locations.oauthClients.create"
`
	// Add some narrowing to reduce noise.
	terms := []string{
		fmt.Sprintf(`protoPayload.resourceName:%q`, c.Name),
		// clientId appears in request/response JSON; try matching text in both.
		fmt.Sprintf(`protoPayload.request:%q`, c.ClientId),
		fmt.Sprintf(`protoPayload.response:%q`, c.ClientId),
	}
	filter = filter + " (" + strings.Join(terms, " OR ") + ")"

	return earliestLogTime(ctx, admin, filter)
}

// findSecretCreateTime returns the earliest Admin Activity timestamp for creating this specific secret.
// We match on the secret's full resource name when possible.
func findSecretCreateTime(ctx context.Context, admin *logadmin.Client, projectID string, c *iam.OauthClient, cr *iam.OauthClientCredential) *time.Time {
	filter := `
resource.type="audited_resource"
protoPayload.serviceName="iam.googleapis.com"
protoPayload.methodName="projects.locations.oauthClients.credentials.create"
`
	// Prefer matching on the credential's full name; also include parent client name as a fallback.
	terms := []string{
		fmt.Sprintf(`protoPayload.response:%q`, cr.Name),
		fmt.Sprintf(`protoPayload.request:%q`, cr.Name),
		fmt.Sprintf(`protoPayload.resourceName:%q`, c.Name),
	}
	filter = filter + " (" + strings.Join(terms, " OR ") + ")"

	return earliestLogTime(ctx, admin, filter)
}

// earliestLogTime scans logs (newest-first by default) and returns the minimum timestamp seen.
func earliestLogTime(ctx context.Context, admin *logadmin.Client, filter string) *time.Time {
	it := admin.Entries(ctx,
		logadmin.Filter(filter),
		// NewestFirst just affects iteration order; we still compute the minimum timestamp.
		logadmin.NewestFirst(),
		// Only Admin Activity contains these, but scoping by filter is enough.
	)
	var min *time.Time
	for {
		e, err := it.Next()
		if err != nil {
			// iterator.Done or real error are both fine; we just stop.
			break
		}
		t := e.Timestamp
		if min == nil || t.Before(*min) {
			cp := t
			min = &cp
		}
	}
	return min
}

func tsOrNA(t *time.Time) string {
	if t == nil {
		return "N/A"
	}
	return t.UTC().Format(time.RFC3339)
}

func short(full string) string {
	// names look like: projects/{p}/locations/global/oauthClients/{id}/credentials/{credId}
	if i := strings.LastIndex(full, "/credentials/"); i >= 0 {
		return full[i+1:]
	}
	return full
}
