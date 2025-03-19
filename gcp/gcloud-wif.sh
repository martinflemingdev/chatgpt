gcloud iam workload-identity-pools providers update-oidc my-wif-provider \
    --workload-identity-pool=my-wif-pool \
    --location=global \
    --attribute-mapping="google.subject=assertion.sub,attribute.aws_role=assertion.arn" \
    --access-token-include-email \
    --access-token-oauth-scopes="https://www.googleapis.com/auth/cloud-platform,https://www.googleapis.com/auth/datacatalog"
