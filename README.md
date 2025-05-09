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

使用golang1.20版本，解决windows7出现的runtime.asmstdcall<0x5e0300060009>问题。

针对常见操作系统的兼容性

* Linux：内核2.6.32及以上
* Windows：Windows7版本及以后，Windows Server 2008及以后
* Darwin：macOS 10.13 High Sierra or 10.14 Mojave.

参考[Go 1.20 Release Notes - The Go Programming Language](https://tip.golang.org/doc/go1.20#windows)


