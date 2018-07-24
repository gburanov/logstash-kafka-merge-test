#!/bin/bash

set -x

/opt/kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic users1 < messages.txt
/opt/kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic users2 < messages.txt
