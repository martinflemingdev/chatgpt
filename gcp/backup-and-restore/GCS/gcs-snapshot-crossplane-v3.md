
# GCP GCS Bucket Snapshot Retention and Restore Strategy (with Crossplane)

## Overview

Google Cloud Storage (GCS) doesn't support native snapshots, but snapshot-like behavior can be achieved through:

- **Object Versioning** ([docs](https://cloud.google.com/storage/docs/object-versioning))
- **Lifecycle Rules** ([docs](https://cloud.google.com/storage/docs/lifecycle))
- **Retention Policies** ([docs](https://cloud.google.com/storage/docs/bucket-lock))
- **Public Access Prevention**
- **Soft Delete Policy** ([docs](https://cloud.google.com/storage/docs/soft-delete))
- **Crossplane Configuration (via `provider-gcp-storage`)**

This document explains how to implement snapshot-like data protection in GCS using Crossplane and GCP features.

---

## 1. Object Versioning

Enables automatic retention of overwritten or deleted object versions.

📚 [Google Cloud Docs: Object Versioning](https://cloud.google.com/storage/docs/object-versioning)

### Crossplane Default:
- `versioning.enabled`: `false` (must be explicitly enabled)

### Crossplane Configuration
```yaml
versioning:
  enabled: true
```

---

## 2. Lifecycle Rules

Manage the lifecycle of GCS objects to control storage costs and automate cleanup.

📚 [Google Cloud Docs: Lifecycle Management](https://cloud.google.com/storage/docs/lifecycle)

### Crossplane Default:
- `lifecycleRules`: none (no automatic cleanup unless defined)

### Example: Delete non-current versions after 30 days
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

Prevents objects from being deleted or overwritten for a specified duration.

📚 [Google Cloud Docs: Bucket Retention Policy](https://cloud.google.com/storage/docs/bucket-lock)

### Crossplane Default:
- `retentionPolicy`: not set (objects can be deleted any time)

### Example: 7-day Retention
```yaml
retentionPolicy:
  retentionPeriod: 604800  # 7 days in seconds
```

---

## 4. Soft Delete Policy

Soft delete retains deleted objects for a defined duration, allowing for recovery even if versioning is disabled.

📚 [Google Cloud Docs: Soft Delete](https://cloud.google.com/storage/docs/soft-delete)

### Crossplane Default:
- `softDeletePolicy`: not set (no soft delete enabled)

### Why Soft Delete Is Relevant to Snapshots

- **Accidental Delete Protection**: Allows easy rollback from deletions.
- **No Versioning Required**: Operates independently from object versioning.
- **Simpler Restore Flow**: No need to track generation numbers.

### Example Crossplane Configuration:
```yaml
softDeletePolicy:
  retentionDurationSeconds: 604800  # 7 days
```

---

## 5. Public Access Prevention

Disables public access to the bucket and all its objects.

📚 [Google Cloud Docs: Public Access Prevention](https://cloud.google.com/storage/docs/public-access-prevention)

### Crossplane Default:
- `publicAccessPrevention`: `inherited` (inherits from project policy)

### Example:
```yaml
publicAccessPrevention: enforced
```

---

## 6. Full Crossplane Bucket YAML

Below is a complete Crossplane manifest for a GCS bucket with versioning, lifecycle, retention policy, soft delete policy, and public access prevention:

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
    location: US
    storageClass: STANDARD
    publicAccessPrevention: enforced
    versioning:
      enabled: true
    softDeletePolicy:
      retentionDurationSeconds: 604800  # 7 days
    retentionPolicy:
      retentionPeriod: 604800  # 7 days in seconds
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

## 7. Limitations of Crossplane for GCS Snapshots

| Feature                          | Supported in Crossplane? | Scope       | Notes |
|----------------------------------|---------------------------|-------------|-------|
| Object Versioning                | ✅ Yes                    | Bucket      | Full control |
| Lifecycle Rules                  | ✅ Yes                    | Bucket      | Fully configurable |
| Retention Policy                 | ✅ Yes                    | Bucket      | Seconds granularity |
| Soft Delete Policy               | ✅ Yes                    | Bucket      | GCP API must support it |
| Public Access Prevention         | ✅ Yes                    | Bucket      | Enforced flag |
| Object Holds (Temp / Legal)     | ❌ No                     | Object      | [Docs](https://cloud.google.com/storage/docs/object-holds) |
| Restore by Generation            | ❌ No                     | Object      | Manual via gsutil |
| Bucket Lock                      | ❌ No                     | Bucket      | Use GCP Console or API |
| IAM Restrictions / Audit Logs    | ✅ Yes (via other CRDs)   | Project     | Use IAM resources |

---

## 8. Summary

Crossplane allows comprehensive configuration of GCS snapshot-like functionality at the **bucket level**, enabling versioning, retention, soft delete, and lifecycle management. Some object-level recovery tasks still require manual interaction via the GCP CLI or console.
