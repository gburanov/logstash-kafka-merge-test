logstash-kafka-merge-test

* Starts kafka and logstash.
* Uses kafka input and output plugins to merge topics
* Also tried to mirror plugins

## Usage
* Get the kafka config
```
/opt/kafka/bin/kafka-topics.sh --describe  --zookeeper zookeeper:2181
```
* Create topics
```
./create_topics.sh
```
* Send messages
```
./send_messages.sh
```
