FROM ubuntu:14.04

MAINTAINER stupendous-man

ADD start-kafka.sh /usr/bin/start-kafka.sh

RUN apt-get update \
	&& apt-get install -y default-jre \
	&& apt-get install -y wget \
	&& mkdir -p ~/Downloads \
	&& wget "http://mirror.cc.columbia.edu/pub/software/apache/kafka/0.10.1.0/kafka_2.11-0.10.1.0.tgz" -O ~/Downloads/kafka.tgz \
	&& mkdir -p ~/kafka \
	&& cd ~/kafka \
	&& tar -xvzf ~/Downloads/kafka.tgz --strip 1 \
	&& sed -r -i "s/(zookeeper.connect)=(.*)/\1=zookeeper:2181/g" ~/kafka/config/server.properties \
	&& chmod a+x /usr/bin/start-kafka.sh

EXPOSE 9092

CMD ["sh", "/usr/bin/start-kafka.sh"]