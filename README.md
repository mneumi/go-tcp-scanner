## TCP 端口扫描器

### 概述

tcp scanner 用于扫描目标机器的端口，查看有哪些端口开放

### 使用方式

```shell
go run main.go [IPv4地址] [开始端口] [结束端口]
```

示例

```shell
go run main.go aaa.bbb.ccc.ddd 1 65535
```
