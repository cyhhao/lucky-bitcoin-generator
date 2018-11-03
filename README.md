# 比特币"靓号"生成工具

## 简介
遵循 bip32 、bip39、 bip44 协议。不停随机生成助记词（私钥）生成地址，筛出满足条件前缀的比特币地址。

例如 1BTChwLPGp9d1aQAgghpU2gzE3zd9PG3Vj 这个有 BTC 前缀的地址就算是靓号啦。

当然你也可以指定长一些的前缀，比如 iLoveBtc，但是越长就越难撞出。

## 使用方式

使用方式有两种，没有 go 环境的可以直接使用编译好的可执行文件。
有开发能力的同学，可以选择自己编译运行。

### 1. 直接使用
根据系统选择执行 main-mac、main-windows.exe、main-linux
```
 cd main/
 ./main-xxx
```

### 2. 自己编译
1. 需要安装dep
2. dep ensure
3. go run main.go

### 输入输出
1. 在执行目录下新建名为 input 的文件
2. input 可写多个要筛选的前缀，可写多行，一行一个
3. 前缀设置的越短，越容易撞出
4. 撞出的结果会收录在 ans 文件夹内

## 注意
base58 编码没有这些字母、数字：

```
大写字母 I
小写字母 l
大写字母 O
数字 0
```

因此靓号地址前缀无法包含以上字母数字，请悉知，不要徒劳。

## 依赖

> [go-bip44](https://github.com/edunuzzi/go-bip44)

> [btcd](https://github.com/btcsuite/btcd)

> [btcutil](https://github.com/btcsuite/btcutil)

> [go-bip39](https://github.com/tyler-smith/go-bip39)
