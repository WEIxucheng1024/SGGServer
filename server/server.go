package main

import (
	"fmt"
	"io"
	"net"
)

var Users  map[string]string

type UserChannle struct {
	FromUser	string
	StrCh		chan string
}

var UserChannles map[string]*UserChannle

func process(conn net.Conn){
	// 循环接收客户端发送的数据
	defer conn.Close()

	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)

		// 1.等待客户端通过conn发送消息
		// 2.如果客户端没有writer[发送],那么协成阻塞在这里
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("客户端退出")
			return
		}

		// 3.显示客户端发送的内容到服务器的终端
		 fmt.Print(string(buf[:n]))


		if string(buf[:10]) == "$userName:" {
			Users[conn.RemoteAddr().String()] = string(buf[11:(n-2)])
		}else if string(buf[:4]) == "$to:" {
			UserChannles[string(buf[5:(n-2)])] = &UserChannle{
				FromUser: Users[string(buf[11:(n-2)])],
				StrCh : make(chan string),
			}
		}

		for k, v := range UserChannles {
			if k == string(buf[11:(n-2)]){
				conn.Write([]byte("返回：" + string(buf[:n])))
				v.
			}
		}

		//服务器返回数据到客户端
		conn.Write([]byte("返回：" + string(buf[:n])))

		fmt.Println(Users)
	}
}

func main() {

	Users = make(map[string]string)
	UserChannles = make(map[string]*UserChannle)

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
		go process(conn)
	}

	fmt.Printf("listen: %v\n",listen)
}
