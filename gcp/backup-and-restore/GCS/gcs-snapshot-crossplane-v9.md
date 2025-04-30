# GCP GCS Bucket Snapshot Retention and Restore Strategy (with Crossplane)

📚 **[Google Cloud Docs: Protection, Backup, and Recovery Overview](https://cloud.google.com/storage/docs/protection-backup-recovery-overview)**

This document describes how to configure snapshot-like behavior and backup retention strategies for GCS buckets using Crossplane and native GCP features.

---

## Key GCS Features for Backup and Restore

| Feature                   | Purpose                                        | Default in GCP / Crossplane | Crossplane Support |
|---------------------------|------------------------------------------------|------------------------------|---------------------|
| Object Versioning         | Preserve prior versions of modified objects   | ❌ Disabled                  | ✅ Yes              |
| Soft Delete Policy        | Retain deleted objects for N days             | ✅ 7 days (604800s)          | ✅ Yes (v1.11+)     |
| Retention Policy          | Prevent deletion before a minimum duration    | ❌ Not set                   | ✅ Yes              |
| Lifecycle Rules           | Auto-delete non-current or aged objects       | ❌ None                      | ✅ Yes              |
| Dual-Region Storage       | Geo-redundancy via predefined or custom pairs | ❌ Defaults to single-region | ✅ Yes              |
| Turbo Replication         | 15-min RPO for dual-region buckets            | ❌ Disabled                  | ❌ No               |
| Cross-Bucket Replication  | Async replication to another GCS bucket       | ❌ Not configured            | ❌ No               |

---

## Object Versioning

📚 [Docs](https://cloud.google.com/storage/docs/object-versioning)

Tracks and preserves older versions of objects when overwritten or deleted.

- **Default**: `false`

### Crossplane Example
```yaml
versioning:
  enabled: true
```

---

## Soft Delete Policy

📚 [Docs](https://cloud.google.com/storage/docs/soft-delete)

Keeps deleted objects in a restorable state for a configurable number of seconds.

- **Default**: `604800` seconds (7 days)

### Custom Example:
```yaml
softDeletePolicy:
  retentionDurationSeconds: 1209600  # 14 days
```

---

## Retention Policy

📚 [Docs](https://cloud.google.com/storage/docs/bucket-lock)

Enforces a minimum retention duration to prevent modification or deletion.

- **Default**: not set

### Crossplane Example
```yaml
retentionPolicy:
  retentionPeriod: 2592000  # 30 days
```

---

## Lifecycle Rules

📚 [Docs](https://cloud.google.com/storage/docs/lifecycle)

Automates the deletion or transition of object versions to manage storage cost and age-based policies.

- **Default**: not configured

### Crossplane Example
```yaml
lifecycleRules:
  - action:
      type: Delete
    condition:
      age: 30
      isLive: false
```

---

## Dual-Region Storage

📚 [Docs](https://cloud.google.com/storage/docs/locations#dual-regions)

Stores data in two distinct regions to provide high availability and resilience.

### Predefined Dual-Region Options

| Dual-Region Name | Region Pair                                 | Description              |
|------------------|----------------------------------------------|--------------------------|
| `ASIA1`          | `asia-northeast1` + `asia-northeast2`        | Tokyo + Osaka            |
| `EUR4`           | `europe-north1` + `europe-west4`             | Finland + Netherlands    |
| `EUR5`           | `europe-west1` + `europe-west2`              | Belgium + London         |
| `EUR7`           | `europe-west2` + `europe-west3`              | London + Frankfurt       |
| `EUR8`           | `europe-west3` + `europe-west6`              | Frankfurt + Zürich       |
| `NAM4`           | `us-central1` + `us-east1`                   | Iowa + South Carolina    |

### Crossplane Example (Predefined)
```yaml
location: nam4
```

### Crossplane Example (Configurable Dual-Region)
```yaml
customPlacementConfig:
  dataLocations:
    - australia-southeast1
    - australia-southeast2
```

---

## Cross-Bucket Replication

📚 [Docs](https://cloud.google.com/storage/docs/replication)

This feature allows **asynchronous replication** of data from one bucket to another, potentially in a different region. It's useful for:
- Creating geographically redundant backups
- Offloading access to regional copies
- Disaster recovery planning

- **Default**: not enabled
- **Crossplane Support**: ❌ Not supported in the GCP storage provider

To configure, use the Console, API, or `gcloud` CLI.

---

## Turbo Replication

📚 [Docs](https://cloud.google.com/storage/docs/turbo-replication)

Turbo Replication is available **only for dual-region buckets**. It ensures that data written to one region is **replicated to the second region within 15 minutes (RPO of 15 mins)**.

Use cases:
- Low RPO requirements (e.g., regulated industries)
- Business continuity across regions

- **Default**: Disabled
- **Crossplane Support**: ❌ Not supported; enable via Console/API

---

## Full Crossplane Bucket YAML Example

```yaml
apiVersion: storage.gcp.upbound.io/v1beta2
kind: Bucket
metadata:
  name: snapshot-enabled-bucket
  labels:
    environment: production
    owner: platform-team
spec:
  forProvider:
    location: nam4
    storageClass: STANDARD
    versioning:
      enabled: true
    softDeletePolicy:
      retentionDurationSeconds: 1209600
    retentionPolicy:
      retentionPeriod: 2592000
    lifecycleRules:
      - action:
          type: Delete
        condition:
          age: 30
          isLive: false
  providerConfigRef:
    name: gcp-provider
```

---

## Restore Strategy Summary

| Scenario                          | Restore Approach                               |
|----------------------------------|------------------------------------------------|
| Soft-deleted object              | Restore via Console or GCS API                 |
| Prior version (versioned)        | Use `gsutil cp gs://bucket/obj#<gen>`          |
| Lifecycle-deleted object         | Not recoverable unless backed up elsewhere     |
| Replicated object (manual setup) | Copy from destination bucket                   |

---

## Summary

Crossplane can declaratively manage most GCS backup features including versioning, soft deletes, retention, lifecycle cleanup, and dual-region replication. Advanced features like Turbo Replication and Cross-Bucket Replication require manual setup using the GCP Console or API.
---

## Google Cloud Storage Transfer Service

📚 [Docs](https://cloud.google.com/storage-transfer/docs/overview)

The **Storage Transfer Service** is a managed solution from Google Cloud that allows you to copy or synchronize data between:

- GCS buckets (intra-cloud)
- AWS S3 or Azure blob storage (inter-cloud)
- On-premises sources via agent

### Use Case for GCS Backups

To create a backup of your bucket on a recurring schedule:

1. Set up **source bucket** (e.g., `gs://prod-data`)
2. Create a **destination bucket** in a different region (e.g., `gs://backup-data`)
3. Use the Storage Transfer Service to copy or synchronize data between them daily/weekly/etc.

### Features

- Supports **daily scheduled syncs**
- Can **delete from destination** to mirror source, or **retain copies**
- Optionally **filters by prefix**, storage class, or modification time
- Can preserve ACLs and metadata

### How to Use (via Console or gcloud)

#### Console
1. Navigate to **Storage Transfer** in the GCP Console
2. Click **Create transfer job**
3. Select **GCS bucket as source**
4. Select **backup bucket as destination**
5. Define scheduling and options

#### CLI Example
```bash
gcloud transfer jobs create gs://prod-data gs://backup-data   --description="Nightly GCS backup"   --schedule-start-date=2025-05-01   --schedule-repeats   --schedule-end-date=2025-12-31   --schedule-time-of-day=02:00
```

> ⚠️ Crossplane does not currently support creating or managing Storage Transfer jobs.

---