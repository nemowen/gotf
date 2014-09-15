package main

import (
	"fmt"
	"net"
	"os"
	//"strconv"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":9999")
	checkErr(err)
	for {
		conn, err := listen.Accept()
		checkErr(err)
		fmt.Println("accept:", conn.RemoteAddr().String())
		go do(conn)
	}
}

func do(conn net.Conn) {
	file, err := os.OpenFile("/Volumes/USB/"+time.Now().Format("060102150405")+".gif", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		conn.Close()
		panic("未打到存储介质:" + err.Error())
		return
	}
	b := make([]byte, 1024)
	defer conn.Close()
	defer file.Close()
	for {
		n, e := conn.Read(b)
		if n > 0 && e == nil {
			stat, _ := file.Stat()
			file.WriteAt(b, stat.Size())
		} else {
			break
		}
	}
	fmt.Println("done!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
