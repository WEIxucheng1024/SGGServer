package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil{
		fmt.Println("dial err :",err)
	}
	fmt.Println("conn 成功，conn = ",conn)
	defer conn.Close()

	for{
		// 1.准备
		// 客户端可以发送单行数据，然后退出
		reader := bufio.NewReader(os.Stdin) // os.stdin 代表标准输入[终端]

		// 从终端读取一行用户输入
		line, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("reader err :",err)
		}

		// 如果用户输入的是exit，那么就退出，这里line的内容是带有\n的，需要切割处理
		trim := strings.Trim(line, "\n")
		if trim == "exit"{
			fmt.Println("客户端退出")
			break
		}

		// 将line发送给服务器
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.writer err = ",err)
		}

		//获取服务器返回的内容
		r := make([]byte,1024)
		_, err = conn.Read(r)
		if err == io.EOF{
			fmt.Println("客户端读取失败")
		}
		fmt.Print(string(r))
	}
}
