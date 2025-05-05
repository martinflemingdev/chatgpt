
# Backup and Restore Options for BigQuery Datasets

This page outlines several strategies and tools for backing up and restoring BigQuery datasets, with a focus on native GCP options and how they integrate with infrastructure-as-code tools like Crossplane.

---

## 🔁 Backup & Recovery Strategy Overview

BigQuery is serverless and highly available by default, but **it does not include automatic point-in-time recovery** for datasets or tables. Backups must be planned and implemented explicitly.

Backup options can be grouped into:

- **Table-level snapshots**
- **Data exports**
- **Dataset copies**
- **Cross-region replication**
- **Scheduled transfers (DTS)**
- **Manual solutions using SQL or API**

---

## 📌 Key BigQuery Backup Options

### 1. Table Snapshots
- Use SQL to create a point-in-time snapshot:
  ```sql
  CREATE SNAPSHOT TABLE dataset.snapshot_table
  CLONE dataset.original_table;
  ```
- **Point-in-time** and copy-on-write.
- **Stored in BigQuery**, billed as storage.
- No automatic scheduling — must be triggered manually or by automation.

> ✅ Useful for low-overhead versioning of critical tables.

---

### 2. Table Copy Jobs
- You can issue a `COPY` job to duplicate tables:
  - Within the same dataset
  - Across datasets
- Can be scripted or managed via API.

---

### 3. Export to Google Cloud Storage (Extract Jobs)
- Export data to GCS in formats like CSV, JSON, AVRO, or PARQUET.
- Supports compression and partitioned files.
- Must be configured per table.

#### Example SQL (for use with `DataTransferConfig`):
```sql
EXPORT DATA OPTIONS(
  uri='gs://my-bucket/backups/table-*.csv',
  format='CSV',
  overwrite=true
) AS
SELECT * FROM my_dataset.my_table;
```

---

### 4. Scheduled Backup with BigQuery Data Transfer Service (DTS)
- **Recommended for automating backups**.
- `DataTransferConfig` can be configured to:
  - Run `EXPORT DATA` SQL jobs on a schedule (daily, hourly, etc.)
  - Copy one dataset to another (cross-region or within project)
- Supported sources:
  - BigQuery-to-BigQuery
  - GCS-to-BigQuery
  - External sources like Ads or Campaign Manager

---

## 📦 Restore Options

### ✅ Restore from Export
- Load data from GCS using `LOAD` jobs or the BigQuery UI.
- Supported formats: CSV, JSON, AVRO, PARQUET.

### ✅ Restore from Table Snapshot
- Use SQL:
  ```sql
  CREATE TABLE my_dataset.restored_table
  CLONE my_dataset.snapshot_table;
  ```

### ✅ Restore from Dataset Copy
- Copy tables back manually.
- Use UI or scripting.

---

## ⚙️ Crossplane Support

### Supported Features

| Feature | Resource | Notes |
|--------|----------|-------|
| Dataset backup to GCS | `Job` or `DataTransferConfig` | Use `EXPORT DATA` SQL or Extract jobs |
| Scheduled backups | `DataTransferConfig` | Can run SQL on a schedule |
| Dataset copy | `DataTransferConfig` | BigQuery-to-BigQuery supported |
| Table snapshot | ❌ Not supported | Must use SQL or API outside Crossplane |

### Example: Scheduled Backup with `DataTransferConfig`

```yaml
apiVersion: bigquery.gcp.upbound.io/v1beta2
kind: DataTransferConfig
metadata:
  name: backup-daily
spec:
  forProvider:
    dataSourceId: scheduled_query
    destinationDatasetId: backup_dataset
    displayName: Daily Export
    location: US
    params:
      query: >
        EXPORT DATA OPTIONS(
          uri='gs://backup-bucket/my_table-*.csv',
          format='CSV',
          overwrite=true
        ) AS SELECT * FROM my_dataset.my_table;
      destination_table_name_template: ""
      write_disposition: WRITE_TRUNCATE
    schedule: every 24 hours
  providerConfigRef:
    name: gcp-provider
```

---

## 🔒 Best Practices

- Enable table-level audit logging.
- Back up critical tables daily via DTS or scheduled jobs.
- For DR, consider **cross-region backups** using `DataTransferConfig`.
- Maintain consistent IAM policies on datasets and GCS buckets used for backup.

---

## 📚 References

- [Google Cloud: BigQuery Backup and DR Strategies](https://cloud.google.com/blog/topics/developers-practitioners/backup-disaster-recovery-strategies-bigquery)
- [Google Cloud KMS for CMEK](https://cloud.google.com/kms/docs/cmek)
- [BigQuery Export Jobs](https://cloud.google.com/bigquery/docs/exporting-data)
