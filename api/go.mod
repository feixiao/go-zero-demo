module github.com/feixiao/go-zero-demo/api

go 1.16

require (
	github.com/feixiao/go-zero-demo/code v0.0.0
	github.com/feixiao/go-zero-demo/rpc v0.0.0
	github.com/tal-tech/go-zero v1.1.8
	google.golang.org/grpc v1.39.0
)

replace (
	github.com/feixiao/go-zero-demo/code v0.0.0 => ../code
	github.com/feixiao/go-zero-demo/rpc v0.0.0 => ../rpc
)
