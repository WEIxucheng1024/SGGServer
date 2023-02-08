package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil{
		fmt.Println("dial err :",err)
	}
	fmt.Println("conn 成功，conn = ",conn)
	defer conn.Close()

	// 1.准备
	// 客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin) // os.stdin 代表标准输入[终端]

	// 从终端读取一行用户输入
	line, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("reader err :",err)
	}
	// 将line发送给服务器
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("conn.writer err = ",err)
	}
	fmt.Printf("客户端发送了 %d个字节",n)
}
