```shell
protoc  --go_out=../../pkg/api --go_opt=paths=source_relative --go-grpc_out=../../pkg/api --go-grpc_opt=paths=source_relative adder.proto

 evans api/proto/adder.proto -p 3200

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


api.Users@127.0.0.1:3200> call ListUsers
offset (TYPE_INT32) => 1
limit (TYPE_INT32) => 1
{}

go run cmd/client/client.go 
Пользователь с email sveta@example.com уже существует
name:"Света"  sex:WOMAN  email:"sveta@example.com"
Пользователь с email serge@example.com не найден
2 [den@example.com sveta@example.com]

```