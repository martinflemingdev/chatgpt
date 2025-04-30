
# GCP GCS Bucket Snapshot Retention and Restore Strategy (with Crossplane)

## Overview

Google Cloud Storage (GCS) doesn't support native snapshots, but snapshot-like behavior can be achieved through:

- **Object Versioning** ([docs](https://cloud.google.com/storage/docs/object-versioning))
- **Lifecycle Rules** ([docs](https://cloud.google.com/storage/docs/lifecycle))
- **Retention Policies** ([docs](https://cloud.google.com/storage/docs/bucket-lock))
- **Public Access Prevention**
- **Soft Delete Policy** ([docs](https://cloud.google.com/storage/docs/soft-delete))
- **Dual-Region Storage** ([docs](https://cloud.google.com/storage/docs/locations#dual-regions))
- **Cross-Bucket Replication** ([docs](https://cloud.google.com/storage/docs/replication))
- **Crossplane Configuration (via `provider-gcp-storage`)**

This document explains how to implement snapshot-like data protection in GCS using Crossplane and GCP features.

---

## 1. Object Versioning

📚 [Docs](https://cloud.google.com/storage/docs/object-versioning)

### Default: `false`

### Crossplane Configuration
```yaml
versioning:
  enabled: true
```

---

## 2. Lifecycle Rules

📚 [Docs](https://cloud.google.com/storage/docs/lifecycle)

### Default: not set

### Crossplane Configuration
```yaml
lifecycleRules:
  - action:
      type: Delete
    condition:
      age: 30
      isLive: false
```

---

## 3. Retention Policy

📚 [Docs](https://cloud.google.com/storage/docs/bucket-lock)

### Default: not set

### Crossplane Configuration
```yaml
retentionPolicy:
  retentionPeriod: 604800  # 7 days
```

---

## 4. Soft Delete Policy

📚 [Docs](https://cloud.google.com/storage/docs/soft-delete)

### Default: not set

### Crossplane Configuration
```yaml
softDeletePolicy:
  retentionDurationSeconds: 604800
```

---

## 5. Public Access Prevention

📚 [Docs](https://cloud.google.com/storage/docs/public-access-prevention)

### Default: `inherited`

### Crossplane Configuration
```yaml
publicAccessPrevention: enforced
```

---

## 6. Dual-Region Storage

📚 [Docs](https://cloud.google.com/storage/docs/locations#dual-regions)

Dual-region buckets store data in two specific geographic locations (e.g., `us-east1` + `us-west1`) to enhance redundancy and availability.

### Default: not used (single-region or multi-region)

### Crossplane Support:
✅ Supported via `location: <dual-region-name>` in bucket spec

```yaml
location: nam4  # Dual region: Iowa + South Carolina
```

---

## 7. Cross-Bucket Replication

📚 [Docs](https://cloud.google.com/storage/docs/replication)

Cross-bucket replication asynchronously copies objects from one bucket to another. Used for geo-redundancy or backup workflows.

### Default: not configured

### Crossplane Support:
❌ Not currently supported directly at the bucket level via Crossplane

Requires manual setup via:
- `gsutil`
- Console
- REST API or Terraform

---

## 8. Turbo Replication

📚 [Docs](https://cloud.google.com/storage/docs/turbo-replication)

Turbo replication is available only for **dual-region buckets** and ensures an RPO (recovery point objective) of 15 minutes for writes to both regions.

### Default: disabled

### Crossplane Support:
❌ Not currently configurable in Crossplane. Must be enabled via GCP Console or API.

---

## 9. Full Crossplane Bucket YAML (with snapshot and retention features)

```yaml
apiVersion: storage.gcp.upbound.io/v1beta2
kind: Bucket
metadata:
  name: bucket-snapshot-hardened
  labels:
    environment: production
    owner: platform-team
spec:
  forProvider:
    location: nam4  # Dual-region for redundancy
    storageClass: STANDARD
    publicAccessPrevention: enforced
    versioning:
      enabled: true
    softDeletePolicy:
      retentionDurationSeconds: 604800
    retentionPolicy:
      retentionPeriod: 604800
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

## 10. Limitations of Crossplane for GCS Snapshots

| Feature                          | Supported in Crossplane? | Scope       | Notes |
|----------------------------------|---------------------------|-------------|-------|
| Object Versioning                | ✅ Yes                    | Bucket      | Full control |
| Lifecycle Rules                  | ✅ Yes                    | Bucket      | Fully configurable |
| Retention Policy                 | ✅ Yes                    | Bucket      | Seconds granularity |
| Soft Delete Policy               | ✅ Yes                    | Bucket      | Requires `v1beta2` CRD |
| Public Access Prevention         | ✅ Yes                    | Bucket      | Enforced flag |
| Dual-Region Storage              | ✅ Yes                    | Bucket      | Use `location: nam4`, etc. |
| Cross-Bucket Replication         | ❌ No                     | Needs Console/API |
| Turbo Replication                | ❌ No                     | Console/API only |
| Object Holds (Temp / Legal)     | ❌ No                     | Object      | [Docs](https://cloud.google.com/storage/docs/object-holds) |
| Restore by Generation            | ❌ No                     | Object      | Manual via gsutil |
| Bucket Lock                      | ❌ No                     | Bucket      | Use GCP Console or API |
| IAM Restrictions / Audit Logs    | ✅ Yes (via other CRDs)   | Project     | Use IAM resources |

---

## 11. Summary

GCP GCS provides several tools for data protection and backup through versioning, retention, soft delete, and regional redundancy. Crossplane supports many of these at the bucket level, while some advanced features like turbo replication and cross-bucket replication must still be configured manually.

