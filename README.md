# RUN DOCKER POSTGRSQL
docker run -d \
  --name pos_postgres \
  --network pos-network \
  -e POSTGRES_USER=posuser \
  -e POSTGRES_PASSWORD=pospass \
  -e POSTGRES_DB=posdb \
  -p 5432:5432 \
  -v ./data:/var/lib/postgresql/data \
  -v ./pos.sql:/docker-entrypoint-initdb.d/pos.sql:ro \
  postgres:15

# ENTRY TO CONTAINER
docker exec -it pos_postgres psql -U posuser -d posdb

# TO RUN PROJECT
go run cmd/main.go

# Setup Auto Reload Project
## Pastikan Air memang terinstall
ls ~/go/bin/air
Kalau tidak ada, install ulang:
go install github.com/air-verse/air@latest

## Tambahkan GOPATH/bin ke PATH (zsh)
1. go env GOPATH
2. nano ~/.zshrc 
3. Add this "export PATH="$PATH:$(go env GOPATH)/bin""
4. source ~/.zshrc
5. air -v
6. air init
7. air

## After air init edit file ".air.toml" copy and paste this:
root = "."
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./tmp/main ./cmd/main.go"
  bin = "tmp/main"
  delay = 1000
  include_ext = ["go", "tpl", "tmpl", "html"]
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_regex = ["_test.go"]
  stop_on_error = true

[color]
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  time = true



