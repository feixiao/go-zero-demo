download:
	@go mod tidy

build:api
	@go build -o cmd

run:download
	@ulimit -HSn 10000
	@go run api.go -f etc/api-api.yaml

api:clean_api
	@goctl api go -api api.api -dir .

clean_api:
	@-rm -rf ./internal/handler/*