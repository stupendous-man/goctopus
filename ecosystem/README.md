# goctopus-ecosystem

The goctopus-ecosystem is being developed as a scripted dockerized environment consisting of multiple containers that can be used to test the functionality of the goctopus framework.

## Current List of Supported Integrations

#### Message Brokers

* Apache Kafka
* NATS

#### Databases

* MongoDB

## Current List of Dependencies

* sarama
* mgo

## Instructions

#### Build goctopus/mongo docker image

```bash
cd mongo/
docker build -t goctopus/mongo .
```

#### Build goctopus/nats docker image

```bash
cd nats/
docker build -t goctopus/nats .
```

#### Build goctopus/kafka docker image

```bash
cd kafka/broker/
docker build -t goctopus/kafka .
```

#### Build goctopus/kafka/producer docker image

```bash
cd kafka/producer/
docker build -t goctopus/kafka/producer .
```

#### Build goctopus/kafka/consumer docker image

```bash
cd kafka/consumer/
docker build -t goctopus/kafka/consumer .
```

#### Initiate goctopus-ecosystem environment

```bash
docker-compose up
```