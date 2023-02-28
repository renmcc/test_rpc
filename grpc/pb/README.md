




protoc -I=./grpc/server/pb --go_out=./grpc/server/pb  --go_opt=module="test-grpc/grpc/pb" ./grpc/server/pb/hello.proto


protoc -I=./grpc/server/pb --go-grpc_out=./grpc/server/pb --go-grpc_opt=module="test-grpc/grpc/pb" ./grpc/server/pb/hello.proto