download:
	@go mod tidy

build:download
	@go build -o cmd

run:
	@go run rpc.go -f etc/rpc.yaml

rpc:
	@goctl rpc proto -src rpc.proto -dir .


model:clean_model
	@goctl model mysql ddl -c -src user.sql -dir ./model

grpcui:
	@grpcui -plaintext localhost:8090

clean_model:
	@rm -rf ./model/*