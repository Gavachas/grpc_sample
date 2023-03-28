Comands
protoc --go_out=. --go_opt=paths=source_relative  grpc_sample/grpc_sample.proto
protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative  grpc_sample/grpc_sample.proto