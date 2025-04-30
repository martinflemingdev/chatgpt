
# GCP GCS Bucket Snapshot Retention and Restore Strategy (with Crossplane)

📚 **[Google Cloud Docs: Protection, Backup, and Recovery Overview](https://cloud.google.com/storage/docs/protection-backup-recovery-overview)**

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

### Default (as of provider-family-storage upgrade):  
```yaml
softDeletePolicy:
  retentionDurationSeconds: 604800
```

Crossplane now sets this by default unless explicitly overridden.

### Why Soft Delete Is Relevant to Snapshots

- **Accidental Delete Protection**: Allows rollback even without versioning.
- **No Versioning Required**
- **Simplified Restore Flow**

### Custom Configuration
```yaml
softDeletePolicy:
  retentionDurationSeconds: 1209600  # 14 days
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

### Default: not used (uses `location: US`)

### Crossplane Configuration:
```yaml
location: nam4  # Dual-region for Iowa + South Carolina
```

---

## 7. Cross-Bucket Replication

📚 [Docs](https://cloud.google.com/storage/docs/replication)

### Default: not enabled

### Crossplane Support:
❌ Not supported directly — must be set via Console, CLI, or API.

---

## 8. Turbo Replication

📚 [Docs](https://cloud.google.com/storage/docs/turbo-replication)

### Default: disabled

### Crossplane Support:
❌ Not supported — enable via GCP Console/API for dual-region buckets only.

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
    location: nam4
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
| Soft Delete Policy               | ✅ Yes (v1.11+)           | Bucket      | Defaults to 7 days |
| Public Access Prevention         | ✅ Yes                    | Bucket      | Enforced flag |
| Dual-Region Storage              | ✅ Yes                    | Bucket      | Use `location` field |
| Cross-Bucket Replication         | ❌ No                     | Needs Console/API |
| Turbo Replication                | ❌ No                     | Console/API only |
| Object Holds (Temp / Legal)     | ❌ No                     | Object      | [Docs](https://cloud.google.com/storage/docs/object-holds) |
| Restore by Generation            | ❌ No                     | Object      | Manual via gsutil |
| Bucket Lock                      | ❌ No                     | Bucket      | Use Console/API |
| IAM Restrictions / Audit Logs    | ✅ Yes (via IAM CRDs)     | Project     | Use Crossplane IAM resources |

---

## 11. Summary

GCP GCS provides tools for resilient storage and recovery using versioning, soft delete, retention, and replication. Crossplane supports most bucket-level features as of `v1.11+`, with defaults such as `softDeletePolicy: 7 days` now visible. Some advanced DR features like Turbo Replication and cross-bucket replication still require external configuration.

