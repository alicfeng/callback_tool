# callback_tool
> Callback debugging terminal tool~



#### 它是什么

`callback_tool` 是一款基于 `Golang` 编写的 `http` 回调调试工具。至于为什么有它的存在，为了方便研发 实施 客户等调试推理平台报警转发与报警回调, 于是 `callback_tool` 就出现了、应用而生。



#### 它支持哪些操作系统

- [x] `Unix`
- [x] `Mac`
- [x] `Windows`



#### 它如何安装

支持**源码编译** 与 **获取可执行的二进制文件**。

- 源码编译

  ```shell
  ➜ git clone https://github.com/alicfeng/callback_tool.git
  ➜ make release
  ➜ ll release
  总用量 13M
  -rwxrwxr-x 1 alicfeng alicfeng 2.3M 7月  13 15:13 callback_tool_mac
  -rw-rw-r-- 1 alicfeng alicfeng 2.2M 7月  13 15:13 callback_tool_mac.tar.gz
  -rwxrwxr-x 1 alicfeng alicfeng 2.2M 7月  13 15:13 callback_tool_unix
  -rw-rw-r-- 1 alicfeng alicfeng 2.2M 7月  13 15:13 callback_tool_unix.tar.gz
  -rwxrwxr-x 1 alicfeng alicfeng 2.1M 7月  13 15:13 callback_tool_win
  -rw-rw-r-- 1 alicfeng alicfeng 2.1M 7月  13 15:13 callback_tool_win.tar.gz
  ```

- 获取可执行的二进制文件

  很懒没有 写脚本 或 `shell`，那就自行从 `release` 下载下来，赋予可执行权限即可



#### 它如何使用

###### 使用文档

```shell
# 帮助文档
➜ callback_tool_unix --help
Usage: callback_tool [options...]
--help  This help text
-h      host.     default 127.0.0.1
-p      port.     default 80
-r      route.    default /api/callback
```

###### 启动调试

```shell
# http://192.168.1.168:8888/api/debugging/callback
callback_tool_unix -h 192.168.1.168 -p 8888 -r /api/debugging/callback
```

> 此时回调工具服务就启动了，处于就绪状态。

###### 调试断言

当在推理平台配置 `http://192.168.1.168:8888/api/debugging/callback` 地址时，当成功接收到请求会将报文内容写进两个文件( **请求体** 与 **客户端信息** )，文件位于程序当前位置。

```shell
➜ ls -a
-rw-rw-r-- 1 alicfeng alicfeng   90 7月  13 15:16 client_20200713_151344_1594624424551275888.ini
-rw-rw-r-- 1 alicfeng alicfeng   18 7月  13 15:13 package_20200713_151344_1594624424551275888.json
```





