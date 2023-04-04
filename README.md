# IntelliDox A File Proecssing System

This is the POC of Intellidox, for users to upload and process files, for example scan virus.

## Functions

1. user upload files
2. user confirm upload
3. user get virus scan results

## System Design

![sytem design diagram](document/images/final.png)

[More detail of the sytem design](document/design.md)

### Frontend: A React app for uploading files and displaying results.
    
    Uploads files directly to S3 using pre-signed URLs.
    Calls the API to confirm the upload and start processing.

### API: RESTful API, web services deployed to EKS
    
    Handles file upload requests and generates pre-signed URLs.
    Handles "confirm" requests and adds messages to the MetadataExtractionQueue.
    Provides document meta data and scan results to clients

### S3: Stores the uploaded files.

### Lambda: Generates pre-signed URLs for S3 uploads.

### EKS Cluster: Contains the following microservices and components:
    
    1. Metadata-extraction service: Extracts metadata from the files.
        Listens to the MetadataExtractionQueue.
        Publishes messages to the MetadataExtractionCompleted SNS topic when the task is completed.
    2. Virus scanner services: Multiple virus scanner pods from different vendors.
        Each pod listens to its own SQS queue (e.g., VirusScannerQueue_A, VirusScannerQueue_B, etc.).
        Each pod's SQS queue is subscribed to the VirusScanningRequested SNS topic.
        Publishes messages to the VirusScanningCompleted SNS topic when the task is completed.
### SNS topics:
    
    1. MetadataExtractionCompleted: Notifies subscribers when the metadata extraction is done.
    2. VirusScanningRequested: Distributes virus scanning tasks to multiple scanner pods.
    3. VirusScanningCompleted: Notifies subscribers when the virus scanning is done.
### SQS queues:
    
    1. MetadataExtractionQueue: Holds tasks for the metadata-extraction service.
    2. VirusScannerQueue_A, VirusScannerQueue_B, etc.: Holds tasks for each virus scanner pod.

### RDS PostgreSQL: Stores file metadata.

### DynamoDB: Stores virus scanning results.

### Monitoring: Datadog for monitoring and alerting.

## 

