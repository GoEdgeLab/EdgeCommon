# 脚本使用指南

## 编译多语言相关源文件
~~~bash
./build-messages.sh
~~~

## 编译API相关源文件
在使用 `build.sh` 编译 `.proto` 文件之前，你需要确保已经为 `protoc` 安装了对应的插件：
~~~bash
# install protoc-gen-go plugin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# install protoc-gen-go-grpc plugin
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
~~~

之后每次 `.proto` 文件有更新的时候，请运行 `build.sh` 重新生成相应的Go源代码和`rpc.json`文件：
~~~bash
./build.sh
~~~

如果文件名有更改，请清空 `pkg/rpc/pb/*.go` 文件，然后再次运行 `build.sh`。


## 生成RPC列表文件
运行：
~~~bash
./proto-json.sh
~~~
可以重新生成 `rpc.json` 文件。