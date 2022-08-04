.PHONY: protoc

protoc:
	protoc -I ./proto --go_out=./ --go-grpc_out=./ proto/awsume.proto
