// 服务器程序实现
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var SavePath string = "/Volumes/USB/" // 文件保存路径
var Suffix string = ".gif"            // 文件后缀

// 主程序入口
func main() {
	// 监听9999端口
	listen, err := net.Listen("tcp", ":9999")
	checkErr(err)
	for {
		conn, err := listen.Accept()
		checkErr(err)
		fmt.Println("accept:", conn.RemoteAddr().String())
		go do(conn)
	}
}

// 可并发的方法
// 为每次连接在指定的目录新建一个文件,循环读取连接中的数据写到文件中.
func do(conn net.Conn) {
	defer conn.Close()
	file, err := os.OpenFile(SavePath+time.Now().Format("060102150405")+Suffix, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		conn.Close()
		panic("未打到存储介质:" + err.Error())
		return
	}
	defer file.Close()
	buffer := make([]byte, 1024)
	for {
		n, e := conn.Read(buffer)
		if n > 0 && e == nil {
			stat, _ := file.Stat()
			file.WriteAt(buffer, stat.Size())
		} else {
			break
		}
	}
	fmt.Println("done!")
}

// 错误检查公共方法
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
