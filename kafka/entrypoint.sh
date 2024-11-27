#!/bin/sh

/etc/confluent/docker/run &

while ! nc -z localhost 9092; do
  echo "Aguardando o Kafka iniciar..."
  sleep 1
done

kafka-topics --create --topic balances --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1 || true

wait