cd go server
go mod init server
goctl api new greet

go mod tidy

under grpc-demo 创建symlink
ln -s $PWD/proto $PWD/go-server/proto

then under $PWD/go-server/
goctl rpc proto -src ./proto/calculator.proto -dir ./rpc