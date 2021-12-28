# cz88

基于 cz88 纯真IP数据库开发的 IP 解析服务 - 支持 http 协议请求或 rpc 协议请求，也支持第三方包的方式引入直接使用

- Go 语言编写
- 进程内缓存结果，重复的 ip 查询响应时间平均为 0.2 ms
- 支持部署后通过 http 协议或 rpc 协议请求服务
- 支持第三方包的方式引入直接使用
- 内置协程数、缓存数、CPU核心数等指标上报（未开放）


## czip.txt 文件获取方式
1. 通过直接安装纯真网络客户端并将其导出为 txt 文件而获得。将其命名为 `czip.txt` 并放置在根目录下即可
2. github releases 下载地址：[czip.zip](https://github.com/NuoMinMin/cz88/releases/download/v1.0.0/czip.zip) 。下载得到 `czip.zip`，解压得到 `czip.txt` 并放置在根目录下即可

## 第三方包引入

    ```go
    package main
     
    import (
        "fmt"
        cz88 "github.com/NuoMinMin/cz88/core"
    )
     
    func main() {
        fmt.Println(cz88.GetIpInfo("210.35.117.200"))
    }
    ```

## 部署方法

### 编译安装
1. 安装 golang 环境。建议 go1.13 以上。
2. 编译运行

    ```shell
    go mod download
    go get github.com/NuoMinMin/cz88
    go build -o ./cz88 github.com/NuoMinMin/cz88/cmd
    ./cz88
    ```

## http 协议请求

| 接口 | 请求方式 | 请求字段 | 说明 |
| :---- | :---- | :---- | :---- |
| /v1/address | GET | ip | 查询 ip 地址相关信息 或 获取访问者自身 ip 地址相关信息（请求字段为空） |
| /v1/my_address | GET |  | 获取访问者自身 ip 地址相关信息 |

## rpc 协议请求

请根据 rpc 客户端语言编译 pb/v1/service.proto 得到 pb 文件 

### go rpc client demo
```
func GetIpInfo() (*v1.AddressResp, error) {
	cc, err := grpc.Dial("127.0.0.1:8108", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	resp, err := v1.NewAppClient(cc).Address(context.TODO(), &v1.AddressReq{
		Ip: "",
	})
	if err != nil {
		errStatus, _ := status.FromError(err)
		if errStatus.Code() == codes.Unavailable {
			return nil, errors.New("No connection could be made because the target machine actively refused it. ")
		}
		return nil, errors.New(errStatus.Message())
	}
	return resp, nil
}
```

## 性能测试

1. 随机 ip 请求测试中，首次响应时长约 50ms/请求，产生缓存后约 0.2 ms/请求
2. 本人开发机压测结果如下所示

```shell
min@Rytia-Envy13
------------------
OS: Ubuntu 20.04.1 LTS on Windows 10 x86_64
Kernel: 4.4.0-19041-Microsoft
Uptime: 24 mins
Packages: 761 (dpkg)
Shell: zsh 5.8
Terminal: /dev/tty2
CPU: Intel i5-8250U (8) @ 1.800GHz
Memory: 6088MiB / 8038MiB

# ab -c 100 -t 10 "http://127.0.0.1:8107/v1/address?ip=210.35.117.200"
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 5000 requests
Completed 10000 requests
Completed 15000 requests
Finished 17113 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8107

Document Path:          /v1/address?ip=210.35.117.200
Document Length:        114 bytes

Concurrency Level:      100
Time taken for tests:   10.041 seconds
Complete requests:      17113
Failed requests:        0
Total transferred:      4073370 bytes
HTML transferred:       1951110 bytes
Requests per second:    1704.25 [#/sec] (mean)
Time per request:       58.677 [ms] (mean)
Time per request:       0.587 [ms] (mean, across all concurrent requests)
Transfer rate:          396.15 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   18  13.8     14      92
Processing:     1   40  20.8     36     197
Waiting:        1   31  17.4     27     180
Total:          1   58  22.8     53     206

Percentage of the requests served within a certain time (ms)
  50%     53
  66%     62
  75%     70
  80%     76
  90%     91
  95%    104
  98%    116
  99%    123
 100%    206 (longest request)
```
