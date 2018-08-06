logstash-kafka-merge-test

* Starts kafka and logstash.
* Uses kafka input and output plugins to merge topics
* Also tried to mirror plugins

## Usage
* Update the hosts file
```
more /etc/hosts
127.0.0.1       localhost kafka
```
* Login to kafka container
```
docker exec -it ae1c92a60f78 bash
```
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
* Check the messages in users topic
```
/opt/kafka/bin/kafka-console-consumer.sh --zookeeper zookeeper:2181 --topic users --from-beginning --property print.key=true --property key.separator="-"
```
* Also in 2 other topics
```
/opt/kafka/bin/kafka-console-consumer.sh --zookeeper zookeeper:2181 --topic users1 --from-beginning --property print.key=true --property key.separator="-"
/opt/kafka/bin/kafka-console-consumer.sh --zookeeper zookeeper:2181 --topic users2 --from-beginning --property print.key=true --property key.separator="-"
```
* Also in replicated topic
```
/opt/kafka/bin/kafka-console-consumer.sh --zookeeper zookeeper:2181 --topic users_replicated --from-beginning
```

## Use recmd
```
kt consume -brokers kafka:9092 -topic users -offsets oldest -encodevalue hex | jq -r .value --unbuffered | recmd proto -p user
```
*
```
kt consume -brokers fh1.dw1.remerge.io:9092 -topic users_eu3 -offsets newest -encodevalue hex | jq -r .value --unbuffered | recmd proto -p user
```
