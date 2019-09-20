proto:
	curl https://raw.githubusercontent.com/hermeschat/proto/master/api.proto > api.proto
	protoc --go_out=plugins=grpc:. api.proto && cp api.pb.go ./api
	rm -rf api.pb.go api.proto
sample-sender:
	go run main.go send 5c4c2683bfd02a2b923af8bf ramzamaliat
sample-receiver:
	go run main.go recv
sample-get-channel:
	go run main.go ch 10a6a18b-e9e8-65d7-7174-99dd794b2dcb
sample-list-channel:
	go run main.go listchannels