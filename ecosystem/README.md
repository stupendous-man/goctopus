# goctupus-ecosystem

The goctupus-ecosystem is being developed as a scripted dockerized environment consisting of multiple containers that can be used to test the functionality of the goctupus framework.

## Current List of Supported Integrations

* Apache Kafka

## Instructions

#### Build goctupus/kafka docker image

```bash
cd kafka/broker/
docker build -t goctupus/kafka .
```

#### Build go/kafka/producer docker image

```bash
cd kafka/producer/
docker build -t go/kafka/producer .
```

#### Build go/kafka/consumer docker image

```bash
cd kafka/consumer/
docker build -t go/kafka/consumer .
```

#### Initiate goctupus-ecosystem environment

```bash
docker-compose up
```