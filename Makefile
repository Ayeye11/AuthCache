# RUN APP
build:
	@go build -o bin/se-thr ./cmd/app

run: build
	@./bin/se-thr

# SQL MIGRATIONS
m := up
migrate:
	@go run ./cmd/migrate $(m)

# GENERATE PROTO FILES (FROM "./internal/router/cache/proto")
gen:
	@protoc -I internal/router/cache/proto --go_out=./internal/router/cache/proto/gen --go_opt=paths=source_relative \
    --go-grpc_out=./internal/router/cache/proto/gen --go-grpc_opt=paths=source_relative \
    internal/router/cache/proto/*.proto

# SCRIPTS
# -- Create migration files:
mfiles:
	@go run ./scripts/migrate $(name)