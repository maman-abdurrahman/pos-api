#!/bin/bash
set -e

CONTAINER_NAME="pos_postgres"
NETWORK_NAME="pos-network"
IMAGE_NAME="pos-api-image"
CONTAINER_API_NAME="pos-api"

# Remove old containers and image
docker rm -f $CONTAINER_API_NAME $CONTAINER_NAME || true
docker rmi -f $IMAGE_NAME || true

# Create Docker network if it doesn't exist
if ! docker network ls --format '{{.Name}}' | grep -w "$NETWORK_NAME" > /dev/null; then
    echo "Creating Docker network: $NETWORK_NAME"
    docker network create $NETWORK_NAME
fi

# Start PostgreSQL
docker run -d \
  --name $CONTAINER_NAME \
  --network $NETWORK_NAME \
  -e POSTGRES_USER=posuser \
  -e POSTGRES_PASSWORD=pospass \
  -e POSTGRES_DB=posdb \
  -p 5432:5432 \
  -v ./data:/var/lib/postgresql/data \
  -v ./pos.sql:/docker-entrypoint-initdb.d/pos.sql:ro \
  postgres:15

# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be ready..."
until docker exec $CONTAINER_NAME pg_isready -U posuser > /dev/null 2>&1; do
    sleep 1
done
echo "PostgreSQL is ready!"

# Build Go app Docker image
docker build -t $IMAGE_NAME .

# Start POS API container
docker run -d \
  --name $CONTAINER_API_NAME \
  --network $NETWORK_NAME \
  -p 9000:9000 \
  $IMAGE_NAME

# Show logs
docker logs -f $CONTAINER_API_NAME
