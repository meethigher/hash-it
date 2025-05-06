# hash-it
轻量级的文件哈希计算工具，支持多种常见的哈希算法

# usage

同步项目依赖

```bash
go mod tidy
```

编译二进制可执行文件

设置环境变量

* GOOS-操作系统
  * windows
  * linux
  * darwin
* GOARCH-架构类型
  * amd64
  * arm64


```bash
# Windows
go build -o hash-it.exe .
# Linux/MacOS
go build -o hash-it .
```


