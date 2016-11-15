# goctupus-ecosystem

The goctupus-ecosystem is being developed as a scripted dockerized environment consisting of a number of containers that can be used to test the functionality of the goctupus framework.

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