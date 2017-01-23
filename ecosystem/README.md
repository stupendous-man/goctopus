# goctupus-ecosystem

The goctupus-ecosystem is being developed as a scripted dockerized environment consisting of multiple containers that can be used to test the functionality of the goctupus framework.

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

#### Build goctupus/mongo docker image

```bash
cd mongo/
docker build -t goctupus/mongo .
```

#### Build goctupus/nats docker image

```bash
cd nats/
docker build -t goctupus/nats .
```

#### Build goctupus/kafka docker image

```bash
cd kafka/broker/
docker build -t goctupus/kafka .
```

#### Build goctupus/kafka/producer docker image

```bash
cd kafka/producer/
docker build -t goctupus/kafka/producer .
```

#### Build goctupus/kafka/consumer docker image

```bash
cd kafka/consumer/
docker build -t goctupus/kafka/consumer .
```

#### Initiate goctupus-ecosystem environment

```bash
docker-compose up
```