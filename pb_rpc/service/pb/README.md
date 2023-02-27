

protoc -I=./pb_rpc/service/pb --go_out=./pb_rpc/service/ --go_opt=module="test-grpc/pb_rpc/service" ./pb_rpc/service/pb/*.proto