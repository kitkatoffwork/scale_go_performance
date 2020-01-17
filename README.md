# scale_go_performance
scale the performance of Go language, comparing with rails and so on ...

## Environment for development
How to build the environment for development of Go language ...
### CentOS7
#### 1. install Go lang
```bash
# 2020/01/18
$ sudo yum install epel-release

$ sudo yum install -y golang
# The volume which will be installed is : 328 M

$ go version
go version go1.13.3 linux/amd64
```

#### 2. setup environment-value for calling Go
```bash
$ mkdir working-directory-path-as-you-like(☆)
$ vim ~/.bash_profile

# ~/.bash_profile
export GOPATH=working-directory-path-as-you-like(☆)
export PATH=$GOPATH/bin:$PATH

$ source ~/.bash_profile
```
#### 3. test Run
```bash
# Run Go file
$ go run filename.go

# Compile Go file
$ go build filename.go
```

> refer to -> [CentOS 7にyumでGoを入れてHello Worldするまで](https://qiita.com/nooboolean/items/11805928527aeb576c21)
