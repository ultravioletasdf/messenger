generate:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.
grpc: 
	grpcui -plaintext localhost:3000
migrate-local:
	atlas schema apply \
	--url sqlite://dev.db \
	--dev-url "sqlite://tmp/dev.db" \
	--to "file://sql/schema.sql"