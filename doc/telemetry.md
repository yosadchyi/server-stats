# Telemetry data handling

There is a set of clients that send telemetry messages in JSON format. 
The general load is 100-200 QPS, with possible spikes up to 5x.

## Task

Build solution using AWS resources to receive, pre-process, and store events. You don't need to write code for handling, focus on designing and explaining the usage of resources.

## Initial analysis

From the question we can infer following qualities of the telemetry data handling system:
- **Scalability** - has HIGH priority both in terms of computation power and storage
- **Availability** - has HIGH priority
- **Durability** - has HIGH priority
- **Latency** - has MEDIUM priority

## Potential solution

To build a solution for receiving, pre-processing,
and storing telemetry messages in JSON format with AWS resources,
we can use a combination of services, high-level architecture follows:

1. **Amazon Kinesis Data Streams:**
    - Kinesis Data Streams can be used to ingest the telemetry messages from clients. 

2. **AWS Lambda:**
    - Lambda function can be used to preprocess the incoming messages from the Kinesis Data Stream
    - Lambda function can be scaled automatically to handle load spikes

3. **Amazon Kinesis Data Firehose:**
    - Lambda function will pass data to a Kinesis Data Firehose delivery stream.

4. **Amazon S3:**
    - Kinesis Data Firehose delivery stream will the pre-processed messages in an Amazon S3 bucket. This provides a durable and scalable storage solution for telemetry data. Data can be organized in the folders.

5. **Amazon CloudWatch:**
    - CloudWatch Alarms can be used to monitor the incoming traffic and trigger autoscaling for Lambda function and/or Kinesis Data Stream when there are spikes in the message volume.

## Outcomes

Proposed architecture covers qualities identified during initial analysis and is able to handle different loads.

## Potential improvements

- DynamoDB can be used to store the pre-processed messages instead of S3. This will allow to query the data in real-time.
- Kinesis Data Analytics can be used to analyze the data in real-time.
- Amazon CloudTrail and AWS Config can be used to audit the system.
- Amazon CloudWatch Logs can be used to monitor the system.
