# [Monolith to Event-Driven Architecture](https://richardchou89.github.io/monolith-to-event-driven-architecture/)

![image](https://github.com/user-attachments/assets/b58d7a8b-d99d-44fd-af45-c4f14ade978d)

This repo contains 2 folders:

1. `/monolith`: A Rails application demonstrating how to send events to SNS.
2. `/microservices`: A consumer who receives events from SNS and process events. A consumers consists of SQS and Lambda.

This repo doesn't contain message broker (SNS). You can use any IaC tools (CDK, Terraform) to create one.
