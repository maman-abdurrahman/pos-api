# RUN DOCKER POSTGRSQL
docker run -d \
  --name pos_postgres \
  -e POSTGRES_USER=posuser \
  -e POSTGRES_PASSWORD=pospass \
  -e POSTGRES_DB=posdb \
  -p 5432:5432 \
  -v ./data:/var/lib/postgresql/data \
  -v ./pos.sql:/docker-entrypoint-initdb.d/pos.sql:ro \
  postgres:15

# ENTRY TO CONTAINER
docker exec -it pos_postgres psql -U posuser -d posdb