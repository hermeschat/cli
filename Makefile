proto:
	curl https://raw.githubusercontent.com/hermeschat/proto/master/api.proto > api.proto
	protoc --go_out=plugins=grpc:. api.proto && cp api.pb.go ./api
	rm -rf api.pb.go api.proto
sample-sender:
	go run main.go send 5c4c2683bfd02a2b923af8bf salamazhermescli
sample-receiver:
	go run main.go recv
sample-get-channel:
	go run main.go ch 7fa90faf-1871-1eb0-acca-e1111f35321b