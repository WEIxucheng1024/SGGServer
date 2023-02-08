package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn){
	// 循环接收客户端发送的数据
	defer conn.Close()

	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)

		fmt.Printf("等待客户端(%s)发送信息",conn.RemoteAddr().String())

		// 1.等待客户端通过conn发送消息
		// 2.如果客户端没有writer[发送],那么协成阻塞在这里
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器端Read err = ", err)
		}

		// 3.显示客户端发送的内容到服务器的终端
		 fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")

	// 1.tcp标识使用的网络协议是tcp
	// 2.0.0.0.0:8888表示监听本地8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err:",err)
		return
	}

	defer listen.Close()	//延时关闭listen

	// 循环等待客户端来连接
	for{
		// 等待客户端来连接
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept() err :",err)
			continue
		}else{
			fmt.Printf("accept() success, comm = %v\n  客户端ip : %v\n", conn, conn.RemoteAddr().String())
		}

		// 这里准备起一个协成，为客户端服务

	}

	fmt.Printf("listen: %v\n",listen)
}
