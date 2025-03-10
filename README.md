# [Monolith to Event-Driven Architecture](https://richardchou89.github.io/monolith-to-event-driven-architecture/)

![image](https://github.com/user-attachments/assets/b58d7a8b-d99d-44fd-af45-c4f14ade978d)

This repo contains 2 folders:

1. `/monolith`: A publisher sending events to SNS.
2. `/microservices`: A consumer receiving events from SNS and processing events. A consumers consists of SQS and Lambda.

This repo doesn't contain message broker (SNS). You can use any IaC tool (CDK, Terraform) to create one.

# Publisher

A Rails application.

# Consumer

An AWS SAM application.
