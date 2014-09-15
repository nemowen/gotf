package main

import (
	//"bytes"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	args := os.Args
	if len(args) < 2 {
		panic("请输入上传文件的绝对路径")
	}
	file, e := os.OpenFile(args[1], os.O_RDONLY, 0666)
	if e != nil {
		panic("文件未找到:" + e.Error())
	}
	defer file.Close()

	b := make([]byte, 1024)
	for {
		num, e := file.Read(b)
		if e == nil && num > 0 {
			_, e := conn.Write(b[0:num])
			if e != nil {
				panic(err)
			}
		} else {
			break
		}
	}

	fmt.Println("上传完毕")

	defer conn.Close()

}
