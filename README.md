# My GRpc



```shell
# 添加依赖
定义proto文件

go get google.golang.org/grpc
go get google.golang.org/protobuf


go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest



protoc --go_out=. --go-grpc_out=. proto/user.proto
```



```shell
protoc -I . \
  -I $(go env GOPATH)/pkg/mod/github.com/googleapis/googleapis@*/ \
  -I $(go list -m -f "{{.Dir}}" github.com/grpc-ecosystem/grpc-gateway/v2) \
  --go_out=. \
  --go-grpc_out=. \
  --grpc-gateway_out=. \
  proto/user.proto

```



# 推荐写法
```shell
option go_package = "github.com/tattoo1880/my-grpc/gen;gen";
```


```shell
--go_out=. \
--go-grpc_out=. \
--grpc-gateway_out=.
```



## 推荐写法
```shell
protoc \
  -I . \
  -I $(go env GOPATH)/pkg/mod/github.com/googleapis/googleapis@*/ \
  --go_out=./gen --go_opt=paths=source_relative \
  --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=./gen --grpc-gateway_opt=paths=source_relative \
  proto/user.proto
  
  
  
  option go_package = "github.com/tattoo1880/myrepo/gen;gen";
```