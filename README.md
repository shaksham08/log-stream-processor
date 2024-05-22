# Log stream processor

1. Ingress (source)

   - We will be supporting multiple source of inputs for this
     - TCP server
     - GRPC server
     - API server

2. Handler

   - Handler will receive the all the logs from the ingress
   - This will process/transform all the logs i.e may be change format or deserialize

3. Filter
   - To filter out very minor logs etc.

- Egress (DB)
  - Sending the log to some Kafka stream , cloud etc

Note:= For now we will be using some dummy source and pass it to Handler

- Now communication between ingress and handler will happen using channel
