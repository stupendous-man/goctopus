FROM ubuntu:14.04

MAINTAINER stupendous-man

RUN mkdir -p /data/db \
	&& apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927 \
	&& echo "deb http://repo.mongodb.org/apt/ubuntu trusty/mongodb-org/3.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.2.list \
	&& apt-get update \
	&& apt-get install -y mongodb-org

EXPOSE 27017

CMD ["mongod"]