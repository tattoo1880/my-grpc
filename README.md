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
protoc -I . \                                                                                                                                                                                    ✔ 
  -I $(go list -m -f "{{.Dir}}" github.com/grpc-ecosystem/grpc-gateway/v2) \
  -I $(go list -m -f "{{.Dir}}" github.com/googleapis/googleapis) \
  --go_out ./gen \
  --go-grpc_out ./gen \
  --grpc-gateway_out ./gen \
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