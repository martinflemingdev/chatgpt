
# GCP GCS Bucket Snapshot Retention and Restore Strategy (with Crossplane)

## Overview

Google Cloud Storage (GCS) doesn't support native snapshots, but snapshot-like behavior can be achieved through:
- **Object Versioning**
- **Lifecycle Rules**
- **Retention Policies**
- **Crossplane Configuration (via `provider-gcp-storage`)**

This document explains how to implement snapshot-like data protection in GCS using Crossplane and GCP features.

---

## 1. Object Versioning

Enables automatic retention of overwritten or deleted object versions.

### Crossplane Configuration
```yaml
versioning:
  enabled: true
```

---

## 2. Lifecycle Rules

Manage the lifecycle of GCS objects to control storage costs and automate cleanup.

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

### Example: 30-day Retention
```yaml
retentionPolicy:
  retentionPeriod: 2592000  # 30 days in seconds
```

---

## 4. Public Access Prevention

Disables public access to the bucket and all its objects.

```yaml
publicAccessPrevention: enforced
```

---

## 5. Full Crossplane Bucket YAML

Below is a complete Crossplane manifest for a GCS bucket with versioning, lifecycle, retention policy, and public access prevention:

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
    location: US
    storageClass: STANDARD
    publicAccessPrevention: enforced
    versioning:
      enabled: true
    retentionPolicy:
      retentionPeriod: 2592000  # 30 days in seconds
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

## 6. Limitations of Crossplane for GCS Snapshots

| Feature                          | Supported in Crossplane? | Scope       | Notes |
|----------------------------------|---------------------------|-------------|-------|
| Object Versioning                | ✅ Yes                    | Bucket      | Full control |
| Lifecycle Rules                  | ✅ Yes                    | Bucket      | Fully configurable |
| Retention Policy                 | ✅ Yes                    | Bucket      | Seconds granularity |
| Public Access Prevention         | ✅ Yes                    | Bucket      | Enforced flag |
| Object Holds (Temp / Legal)     | ❌ No                     | Object      | Use GCP CLI/API |
| Restore by Generation            | ❌ No                     | Object      | Manual via gsutil |
| Bucket Lock                      | ❌ No                     | Bucket      | Requires GCP Console or API |
| IAM Restrictions / Audit Logs    | ✅ Yes (via other CRDs)   | Project     | Use IAM resources |

---

## 7. Summary

Crossplane allows comprehensive configuration of GCS snapshot-like functionality at the **bucket level**, enabling versioning, retention, and lifecycle management. Some object-level recovery tasks still require manual interaction via the GCP CLI or console.

