#!/bin/bash
set -e

# CONTAINER_NAME="pos_postgres"
# NETWORK_NAME="pos-network"
# IMAGE_NAME="pos-api-image"
# CONTAINER_API_NAME="pos-api"

# # Remove old containers and image
# docker rm -f $CONTAINER_API_NAME $CONTAINER_NAME || true
# docker rmi -f $IMAGE_NAME || true

# # Create Docker network if it doesn't exist
# if ! docker network ls --format '{{.Name}}' | grep -w "$NETWORK_NAME" > /dev/null; then
#     echo "Creating Docker network: $NETWORK_NAME"
#     docker network create $NETWORK_NAME
# fi

# # Start PostgreSQL
# docker run -d \
#   --name $CONTAINER_NAME \
#   --network $NETWORK_NAME \
#   -e POSTGRES_USER=posuser \
#   -e POSTGRES_PASSWORD=pospass \
#   -e POSTGRES_DB=posdb \
#   -p 5432:5432 \
#   -v ./data:/var/lib/postgresql/data \
#   -v ./pos.sql:/docker-entrypoint-initdb.d/pos.sql:ro \
#   postgres:15


# # Build Go app Docker image
# docker build -t $IMAGE_NAME .

# # Start POS API container
# docker run -d \
#   --name $CONTAINER_API_NAME \
#   --network $NETWORK_NAME \
#   -p 9000:9000 \
#   $IMAGE_NAME

# # Show logs
# docker logs $CONTAINER_API_NAME


# Container and network names
POSTGRES_CONTAINER="pos_postgres"
REDIS_CONTAINER="pos_redis"
RABBITMQ_CONTAINER="pos_rabbitmq"
NETWORK_NAME="pos-network"
IMAGE_NAME="pos-api-image"
API_CONTAINER="pos-api"

# Remove old containers and image
docker rm -f $API_CONTAINER $POSTGRES_CONTAINER $REDIS_CONTAINER $RABBITMQ_CONTAINER || true
docker rmi -f $IMAGE_NAME || true

# Create Docker network if it doesn't exist
if ! docker network ls --format '{{.Name}}' | grep -w "$NETWORK_NAME" > /dev/null; then
    echo "Creating Docker network: $NETWORK_NAME"
    docker network create $NETWORK_NAME
fi

# Start PostgreSQL
docker run -d \
  --name $POSTGRES_CONTAINER \
  --network $NETWORK_NAME \
  -e POSTGRES_USER=posuser \
  -e POSTGRES_PASSWORD=pospass \
  -e POSTGRES_DB=posdb \
  -p 5432:5432 \
  -v ./data:/var/lib/postgresql/data \
  -v ./pos.sql:/docker-entrypoint-initdb.d/pos.sql:ro \
  postgres:15

# Start Redis
docker run -d \
  --name $REDIS_CONTAINER \
  --network $NETWORK_NAME \
  -p 6379:6379 \
  redis:7

# Start RabbitMQ (with management UI)
docker run -d \
  --name $RABBITMQ_CONTAINER \
  --network $NETWORK_NAME \
  -p 5672:5672 \    # AMQP protocol
  -p 15672:15672 \  # Management UI
  -e RABBITMQ_DEFAULT_USER=guest \
  -e RABBITMQ_DEFAULT_PASS=guest \
  rabbitmq:3-management

# Build Go app Docker image
docker build -t $IMAGE_NAME .

# Start POS API container
docker run -d \
  --name $API_CONTAINER \
  --network $NETWORK_NAME \
  -p 9000:9000 \
  $IMAGE_NAME

# Show logs of API container
docker logs $API_CONTAINER
