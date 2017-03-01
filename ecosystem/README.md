# groketa-ecosystem

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

#### Build groketa/mongo docker image

```bash
cd mongo/
docker build -t groketa/mongo .
```

#### Build groketa/nats docker image

```bash
cd nats/
docker build -t groketa/nats .
```

#### Build groketa/kafka docker image

```bash
cd kafka/broker/
docker build -t groketa/kafka .
```

#### Build groketa/kafka/producer docker image

```bash
cd kafka/producer/
docker build -t groketa/kafka/producer .
```

#### Build groketa/kafka/consumer docker image

```bash
cd kafka/consumer/
docker build -t groketa/kafka/consumer .
```

#### Initiate groketa-ecosystem environment

```bash
docker-compose up
```