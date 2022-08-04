.PHONY: protoc build tidy

protoc:
	protoc -I ./proto --go_out=./ --go-grpc_out=./ proto/awsume.proto
build:
	go build -o awsume
tidy:
	go mod tidy
