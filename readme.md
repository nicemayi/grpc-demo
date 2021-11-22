cd go server
go mod init server
goctl api new greet

go mod tidy

under grpc-demo 创建symlink
ln -s $PWD/proto $PWD/go_server/proto

under go_server
protoc --go_out=. --go-grpc_out=. proto/go_server/calculator.proto

protoc --python_out=. --python_grpc_out=. proto/go_server/calculator.proto
ln -s $PWD/proto $PWD/python_client/proto


protoc --go_out=. --go-grpc_out=. proto/go_server/calculator.proto