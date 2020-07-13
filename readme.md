## 回调调试工具 ( CallbackTool )

#### 前言概览

为了方便研发 实施 客户等调试推理平台报警转发与报警回调, 于是 `callback_tool` 就出现了、应用而生。它是一款基于 go 语言编写的一个命令行工具，适用于 Linux、Mac、Windows等平台。
那么它可以做什么? 它只有一个功能:接收请求并将请求报文写入文件。

#### 安装使用
###### 安装

> 支持 `Mac` `Unix` `Win` 

http://git.extremevision.com.cn/yumen/callback_tool/tags/1.0.0

###### 使用
```shell
# 帮助文档
➜ callback_tool_unix --help
Usage: callback_tool [options...]
--help  This help text
-h      host.     default 127.0.0.1
-p      port.     default 80
-r      route.    default /api/callback

# 启动服务 | http://192.168.1.168:8888/api/callback
nobup callback_tool_unix -h 192.168.1.168 -p 8888 -r /api/callback &
```

此时回调工具服务就启动了，处于就绪状态。

当在推理平台配置 `http://192.168.1.168:8888/api/callback` 地址时，当成功接收到请求会将报文内容写进一个文件，文件位于程序当前位置。

```shell
➜  ll
总用量 2.2M
-rwxrwxr-x 1 alicfeng alicfeng 2.2M 7月  10 12:16 callback_tool
-rw-rw-r-- 1 alicfeng alicfeng   18 7月  10 12:17 package_20200710_121705_1594354625644399968.json
```