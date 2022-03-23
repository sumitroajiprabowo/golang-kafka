## Golang with Kafka using Docker

### Run Dokcer
```
docker-compose up
```

### List Topic Kafka
Download Binary Kafka https://kafka.apache.org/downloads
```
sh kafka-topics.sh --bootstrap-server=localhost:29092 --list
```

### Create Topic
```
sh kafka-topics.sh --create --bootstrap-server=localhost:29092 --replication-factor=1 --partitions=1 --topic=test
```

### Kafka Consumer
```
sh kafka-console-consumer.sh --bootstrap-server localhost:29092 --topic test --from-beginning
```
Create consumer again
```
sh kafka-console-consumer.sh --bootstrap-server localhost:29092 --topic testagain --from-beginning
```

### Kafka Producer
```
sh kafka-console-producer.sh --broker-list localhost:29092 --topic test
```

### Kafka Group
```
sh kafka-console-consumer.sh --bootstrap-server localhost:29092 --topic test --group testgroup
```


