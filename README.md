## Go-MySQL-Pools

 一个简单、跨平台的MySQL连接池Server工具。

### Requirements

- Go 1.7 +
- MySQL(4.1+)
- Go-MySQL-Driver

### How to Install

```shell
$ go get -u github.com/go-sql-driver/mysql
$ go get -u github.com/GraySilver/go_mysql_pools
```

进入对应gopath路径，进入GraySilver/go_mysql_pools/src目录下。

执行以下命令：

```shell
$ go build
$ go install
```

即可在gopath/下bin目录中找到go_mysql_pools编译完成的执行程序。

### Quick Start

该示例在Win10，Golang1.9下进行。

```shell
$ src.exe -h
```

得到以下参数：

```shell
Usage of src.exe:
  -c string  // MySQL连接参数，遵循Go-MySQL-Driver Engine写法
        MySQL connect engine.
  -h string
        http listen host. (default "127.0.0.1") // 对外HOST地址
  -maxIdleConns int
        MySQL maxIdleConns (default 5)	// MySQL连接池最大闲置数
  -maxOpenConns int
        MySQL maxOpenConns (default 5) 	// MySQL连接池最大连接数
  -p string
        http listen port. (default "8080")	// 对外Port端口
```

运行以上程序，若成功则显示以下信息：

```shell
2018/06/17 18:36:27 Conecting...
2018/06/17 18:36:27 Connecting suc. MaxOpenConns: 5 , MaxIdleConns: 5
```

