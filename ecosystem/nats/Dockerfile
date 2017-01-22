FROM ubuntu:14.04

MAINTAINER stupendous-man

ADD start-nats.sh /usr/bin/start-nats.sh

RUN apt-get update \
	&& apt-get install -y wget \
	&& apt-get install -y unzip \
	&& mkdir -p ~/Downloads \
	&& wget "https://github.com/nats-io/gnatsd/releases/download/v0.9.4/gnatsd-v0.9.4-linux-amd64.zip" -O ~/Downloads/gnatsd-v0.9.4-linux-amd64.zip \
	&& unzip -p ~/Downloads/gnatsd-v0.9.4-linux-amd64.zip gnatsd-v0.9.4-linux-amd64/gnatsd > gnatsd \
	&& chmod +x gnatsd \
	&& mv gnatsd /usr/bin/gnatsd

EXPOSE 4222

CMD ["sh", "/usr/bin/start-nats.sh"]