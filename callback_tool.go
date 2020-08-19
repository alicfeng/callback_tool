/*
Author   :    AlicFeng
Email    :    a@samego.com
Github   :    https://github.com/alicfeng/callback_tool.git
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
	output  = flag.String("o", ".", "output directory")
	version = flag.Bool("v", false, "show version and exit")
)

const (
	DirectorySeparator = "/"
	FileSignPackage    = "_package.json"
	FileSignClient     = "_client.ini"
)

type reqStructure struct {
	ip     string
	port   int
	url    string
	method string
	body   string
}

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
			"-o      output.   default ." + "\n" +
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
	// 1.参数处理
	body, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	ip, port, _ := net.SplitHostPort(request.RemoteAddr)
	req := reqStructure{}
	req.port, _ = strconv.Atoi(port)
	req.ip = ip
	req.body = string(body)
	req.url = request.URL.Host + request.URL.Path

	// 2.异步处理回调业务
	go callbackService(req)

	// 3.响应客户端请求
	_, _ = io.WriteString(rw, "successful")
}

/**
接收报文并写进文件Service
*/
func callbackService(request reqStructure) {
	// 1.创建报文文件 including package as well as client
	var packageFile, clientFile *os.File
	t := time.Now()
	date := t.Format("20060102")
	directory := *output + DirectorySeparator + date
	prefix := date + DirectorySeparator + t.Format("15.04.05") + "_" + strconv.FormatInt(t.UnixNano(), 10)
	if _, err := os.Stat(directory); err != nil {
		_ = os.MkdirAll(directory, 0777)
	}

	// 2.记录请求报文信息
	packageFile, err := os.Create(*output + DirectorySeparator + prefix + FileSignPackage)
	if nil != err {
		fmt.Println("starting ???")
		panic(err)
		return
	}
	defer packageFile.Close()

	_, _ = io.WriteString(packageFile, request.body)

	// 3.记录客户端信息
	clientFile, err = os.Create(*output + DirectorySeparator + prefix + FileSignClient)
	if nil != err {
		panic(err)
		return
	}
	defer clientFile.Close()

	_, _ = io.WriteString(clientFile, "time="+prefix+
		"\nip = "+request.ip+
		"\nport = "+strconv.Itoa(request.port)+
		"\nmethod = "+request.method+
		"\nurl = "+request.url)
}
