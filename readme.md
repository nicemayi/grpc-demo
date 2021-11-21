cd go server
go mod init server
goctl api new greet

go mod tidy

ln -s ./proto/ ./go-server/calculator/rpc/
