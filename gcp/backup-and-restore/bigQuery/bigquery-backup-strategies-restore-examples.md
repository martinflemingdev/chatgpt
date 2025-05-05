
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
```sql
CREATE SNAPSHOT TABLE dataset.snapshot_table
CLONE dataset.original_table;
```

#### ✅ Crossplane Example (via `DataTransferConfig`)
```yaml
apiVersion: bigquery.gcp.upbound.io/v1beta2
kind: DataTransferConfig
metadata:
  name: snapshot-table-daily
spec:
  forProvider:
    project: non-app-1234
    location: us-east1
    dataSourceId: scheduled_query
    destinationDatasetIdSelector:
      matchLabels:
        usage: snapshot-target
    displayName: Daily Table Snapshot
    params:
      query: >
        CREATE SNAPSHOT TABLE my_dataset.snapshot_table
        CLONE my_dataset.original_table;
      destination_table_name_template: ""
      write_disposition: WRITE_EMPTY
    schedule: "every 24 hours"
```

---

### 2. Table Copy Jobs
```sql
CREATE TABLE backup_dataset.my_table_copy
AS SELECT * FROM my_dataset.my_table;
```

#### ✅ Crossplane Example
```yaml
apiVersion: bigquery.gcp.upbound.io/v1beta2
kind: DataTransferConfig
metadata:
  name: copy-table-daily
spec:
  forProvider:
    project: non-app-1234
    location: us-east1
    dataSourceId: scheduled_query
    destinationDatasetIdSelector:
      matchLabels:
        usage: table-copy-target
    displayName: Daily Table Copy
    params:
      query: >
        CREATE TABLE backup_dataset.my_table_copy
        AS SELECT * FROM my_dataset.my_table;
      destination_table_name_template: ""
      write_disposition: WRITE_TRUNCATE
    schedule: "every 24 hours"
```

---

### 3. Export to Google Cloud Storage
```sql
EXPORT DATA OPTIONS(
  uri='gs://my-bucket/backups/table-*.csv',
  format='CSV',
  overwrite=true
) AS
SELECT * FROM my_dataset.my_table;
```

#### ✅ Crossplane Example
```yaml
apiVersion: bigquery.gcp.upbound.io/v1beta2
kind: DataTransferConfig
metadata:
  name: extract-to-gcs-daily
spec:
  forProvider:
    project: non-app-1234
    location: us-east1
    dataSourceId: scheduled_query
    destinationDatasetIdSelector:
      matchLabels:
        usage: extract-log-target
    displayName: Daily Table Export
    params:
      query: >
        EXPORT DATA OPTIONS(
          uri='gs://my-bucket/backups/my_table-*.csv',
          format='CSV',
          overwrite=true
        ) AS
        SELECT * FROM my_dataset.my_table;
      destination_table_name_template: ""
      write_disposition: WRITE_EMPTY
    schedule: "every 24 hours"
```

---

## 📦 Restore Options

### ✅ Restore from Export
Re-import exported files from GCS into BigQuery using `LOAD DATA`.

```sql
LOAD DATA INTO my_dataset.my_table
FROM FILES (
  format = 'CSV',
  uris = ['gs://my-bucket/backups/my_table-0000.csv']
);
```

#### ✅ Crossplane Example
```yaml
apiVersion: bigquery.gcp.upbound.io/v1beta2
kind: Job
metadata:
  name: load-from-export
spec:
  forProvider:
    project: non-app-1234
    location: us-east1
    jobReference:
      jobId: load-from-gcs-job
    configuration:
      load:
        destinationTable:
          datasetId: my_dataset
          projectId: non-app-1234
          tableId: my_table
        sourceUris:
          - gs://my-bucket/backups/my_table-0000.csv
        sourceFormat: CSV
        autodetect: true
        writeDisposition: WRITE_APPEND
```

---

### ✅ Restore from Table Snapshot
```sql
CREATE TABLE my_dataset.restored_table
CLONE my_dataset.snapshot_table;
```

#### ✅ Crossplane Example
```yaml
apiVersion: bigquery.gcp.upbound.io/v1beta2
kind: DataTransferConfig
metadata:
  name: restore-from-snapshot
spec:
  forProvider:
    project: non-app-1234
    location: us-east1
    dataSourceId: scheduled_query
    destinationDatasetIdSelector:
      matchLabels:
        usage: snapshot-restore
    displayName: Restore from Snapshot
    params:
      query: >
        CREATE TABLE my_dataset.restored_table
        CLONE my_dataset.snapshot_table;
      destination_table_name_template: ""
      write_disposition: WRITE_EMPTY
    schedule: "every 24 hours"
```

---

### ✅ Restore from Dataset Copy
Manually copy all backed up tables into the original dataset.

#### ✅ Crossplane Example
```yaml
apiVersion: bigquery.gcp.upbound.io/v1beta2
kind: DataTransferConfig
metadata:
  name: restore-dataset-copy
spec:
  forProvider:
    project: non-app-1234
    location: us-east1
    dataSourceId: scheduled_query
    destinationDatasetIdSelector:
      matchLabels:
        usage: dataset-restore
    displayName: Restore Dataset Copy
    params:
      query: >
        CREATE OR REPLACE TABLE original_dataset.my_table
        AS SELECT * FROM backup_dataset.my_table;
      destination_table_name_template: ""
      write_disposition: WRITE_TRUNCATE
    schedule: "every 24 hours"
```

---

## ⚙️ Crossplane Support Summary

| Feature                     | Resource            | Notes |
|----------------------------|---------------------|-------|
| Dataset backup to GCS      | `DataTransferConfig`| Use `EXPORT DATA` SQL |
| Scheduled backups           | `DataTransferConfig`| Can run SQL on a schedule |
| Dataset copy               | `DataTransferConfig`| BigQuery-to-BigQuery supported |
| Table snapshot (via SQL)   | `DataTransferConfig`| SQL-based snapshot possible |
| Ad-hoc job execution       | `Job`               | For one-time query/extract/load |
| Restore from GCS           | `Job`               | Load CSV, JSON, etc. from GCS |
| Restore from snapshot      | `DataTransferConfig`| SQL-based `CLONE` |
| Restore dataset copy       | `DataTransferConfig`| Recreates original tables |

---

## 🔒 Best Practices

- Enable table-level audit logging.
- Back up critical tables daily via DTS or scheduled queries.
- For DR, consider **cross-region backups** using `DataTransferConfig`.
- Maintain consistent IAM policies on datasets and GCS buckets used for backup.

---

## 📚 References

- [Google Cloud: BigQuery Backup and DR Strategies](https://cloud.google.com/blog/topics/developers-practitioners/backup-disaster-recovery-strategies-bigquery)
- [BigQuery Export Jobs](https://cloud.google.com/bigquery/docs/exporting-data)
- [BigQuery SQL Reference: EXPORT DATA](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#export_data_statement)
