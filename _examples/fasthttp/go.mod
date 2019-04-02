module github.com/karmi/go-elasticsearch/_examples/fasthttp

go 1.11

replace github.com/karmi/go-elasticsearch => ../..

require (
	github.com/karmi/go-elasticsearch master
	github.com/valyala/fasthttp v1.1.0
)
