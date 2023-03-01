




protoc -I=./grpc/pb --go_out=./grpc/pb  --go_opt=module="test-grpc/grpc/pb" ./grpc/pb/hello.proto


protoc -I=./grpc/pb --go-grpc_out=./grpc/pb --go-grpc_opt=module="test-grpc/grpc/pb" ./grpc/pb/hello.proto