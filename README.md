# 比特币"靓号"生成工具

## 简介
遵循 bip32 、bip39、 bip44 随机生成比特币地址和相应的助记词，筛选出满足条件前缀的『靓号地址』

## 使用方式

### 直接使用
根据系统选择执行 main-mac、main-windows.exe、main-linux
```
 cd main/
 ./main-xxx
```

## 编译
1. 需要安装dep
2. dep ensure
3. go run main.go

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