g:
	protoc --gofast_out=plugins=grpc:. grpc/proto/user.proto
s:
	go run grpc/server/server.go
c:
	go run grpc/client/client.go