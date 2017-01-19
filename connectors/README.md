# goctopus-connectors

The goctopus-connectors package is a pure Go client library for interfacing with various messaging systems with a focus on simplicity. A high level API is provided through the `Connector` struct.

## Current List of Supported Integrations

* Apache Kafka
* [NATS](http://nats.io/documentation/)  
Lightweight cloud-native messaging system written in Go. Provides among other features:
    * Pure pub-sub
    * Auto pruning of subscribers
    * Text based protocol (can Telnet)
    * Event streaming
* [NSQ](http://nsq.io/overview/quick_start.html)
Distributed and decentralized topology with no SPOF, horizotaly scalable, primarily in-memory