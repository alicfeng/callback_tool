/*
Author   :    AlicFeng
Email    :    a@samego.com
Gitlab   :    http://git.extremevision.com.cn/yumen/callback_tool.git
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

/**
cli parameters
*/
var (
	host    = flag.String("h", "127.0.0.1", "host(127.0.0.1)")
	port    = flag.Int("p", 80, "port(80)")
	route   = flag.String("r", "/api/callback", "route path(/api/callback)")
	version = flag.Bool("v", false, "show version and exit")
)

/**
init func
*/
func init() {
	// init flag for command
	flag.CommandLine.Usage = func() {
		fmt.Println("Usage: callback_tool [options...]\n" +
			"--help  This help text" + "\n" +
			"-h      host.     default 127.0.0.1" + "\n" +
			"-p      port.     default 80" + "\n" +
			"-r      route.    default /api/callback" +
			"")
		os.Exit(0)
	}
	flag.Parse()
	if *version {
		fmt.Println("callback_tool version: 1.0.0")
		os.Exit(0)
	}
}

/**
main func
*/
func main() {
	http.HandleFunc(*route, callbackHandle)
	_ = http.ListenAndServe(*host+":"+strconv.Itoa(*port), nil)
}

/**
接收报文并写进文件Handle
*/
func callbackHandle(rw http.ResponseWriter, request *http.Request) {
	// 1.创建报文文件 including package as well as client
	var packageFile, clientFile *os.File
	t := time.Now()
	prefix := t.Format("20060102_150405") + "_" + strconv.FormatInt(t.UnixNano(), 10)

	// 1.1记录请求报文信息
	packageOutput := "package_" + prefix + ".json"
	packageFile, err := os.Create(packageOutput)
	if nil != err {
		panic(err)
		return
	}

	body, _ := ioutil.ReadAll(request.Body)
	_, _ = io.WriteString(packageFile, string(body))

	// 1.2记录客户端信息
	clientOutput := "client_" + prefix + ".ini"
	ip, port, err := net.SplitHostPort(request.RemoteAddr)
	clientFile, err = os.Create(clientOutput)
	if nil != err {
		panic(err)
		return
	}
	_, _ = io.WriteString(clientFile, "time="+prefix+
		"\nip = "+ip+
		"\nport = "+port+
		"\nmethod = "+request.Method+
		"\nurl = "+request.URL.Host+request.URL.Path)

	// 2.响应客户端请求
	_, _ = io.WriteString(rw, "successful")
}
