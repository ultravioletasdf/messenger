generate:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.
grpc: 
	grpcui -plaintext localhost:3000
migrate-local:
	atlas schema apply \
	--url sqlite://dev.db \
	--dev-url "sqlite://tmp/dev.db" \
	--to "file://sql/schema.sql"
package:
	npx grpc_tools_node_protoc \
	--ts_proto_out=../pkg/node \
	--ts_proto_opt=outputServices=nice-grpc,outputServices=generic-definitions,useExactTypes=false,esModuleInterop=true,importSuffix=.js \
	--proto_path=proto \
	proto/*.proto