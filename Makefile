dep:
	go mod tidy
	go mod vendor        

proto:
	protoc --proto_path=pb pb/*.proto --go_out=gen/ 
	protoc --proto_path=pb pb/*.proto --go-grpc_out=require_unimplemented_servers=false:gen/
	protoc -I pb --grpc-gateway_out ./gen/  \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	pb/*.proto