# grpc-jaeger

[![Build Status](https://travis-ci.org/moxiaomomo/grpc-jaeger.svg?branch=master)](https://travis-ci.org/moxiaomomo/grpc-jaeger)
[![Go Report Card](https://goreportcard.com/badge/github.com/moxiaomomo/grpc-jaeger)](https://goreportcard.com/badge/github.com/moxiaomomo/grpc-jaeger)

grpc-jaeger是基于Go的针对gRPC的一种拦截器实现，用于结合jaeger来实现rpc调用链跟踪。

# Dependencies

```
github.com/opentracing/opentracing-go
github.com/uber/jaeger-client-go
```

# Example

## 1）部署jaeger

假设在某节点已部署好jaeger服务，<br>
jaeger-agent地址为`192.168.1.100:6831`，<br>
jaeger-ui地址为`http://192.168.1.100:16686/`。

## 2）运行测试

测试用例在wrapper_test.go中，直接在该文件当前目录中运行`go test`

```bash
pintai@MG:/yourpath/grpc-jaeger$ go test
2018/06/17 15:06:05 Initializing logging reporter
2018/06/17 15:06:06 Initializing logging reporter
SayHello Called.
2018/06/17 15:06:06 Reporting span 107bf923fcfc238e:1f607766f1329efd:107bf923fcfc238e:1
2018/06/17 15:06:06 Reporting span 107bf923fcfc238e:107bf923fcfc238e:0:1
call sayhello suc, res:message:"Hi im tester\n"
PASS
ok  	github.com/moxiaomomo/grpc-jaeger	3.004s
```

## 3）查看jaegerUI
打开`http://192.168.1.100:16686/`后查询对应Service， 可看到以下跟踪结果：
![jaegerui](./jaegerui.png)