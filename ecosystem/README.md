# goctupus-ecosystem

The goctupus-ecosystem is being developed as a scripted dockerized environment consisting of a number of containers that can be used to test the functionality of the goctupus framework.

## Instructions

### Build goctupus/kafka docker image

```docker
docker build -t goctupus/kafka .
```

### Build go/kafka/producer docker image

```docker
docker build -t go/kafka/producer .
```

### Build go/kafka/consumer docker image

```docker
docker build -t go/kafka/consumer .
```

### Initiate goctupus-ecosystem environment

```docker
docker-compose up
```