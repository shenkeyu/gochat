package fuwuqi

import (
	"fmt"
	"net"
)

////////////////////////////////////////////////////////
//
//服务器端接收数据线程
//参数：
//      数据连接 conn
//      通讯通道 messages
//
////////////////////////////////////////////////////////
func FwqReceive(conn net.Conn, messages chan string) {

	fmt.Println("连接来自：", conn.RemoteAddr().String())
	buf := make([]byte, 1024)
	for {
		lenght, err := conn.Read(buf)
		if checkError(err, "接收数据信息：") == false {
			conn.Close()
			break
		}
		if lenght > 0 {
			buf[lenght] = 0
		}
		reciveStr := string(buf[0:lenght])
		messages <- reciveStr

	}

}

////////////////////////////////////////////////////////
//
//服务器发送数据的线程
//
//参数
//      连接字典 conns
//      数据通道 messages
//
////////////////////////////////////////////////////////
func FwqSend(conns *map[string]net.Conn, messages chan string) {

	for {
		msg := <-messages
		fmt.Println(msg)

		for key, value := range *conns {

			fmt.Println("数据发送来自：", key)
			_, err := value.Write([]byte(msg))
			if err != nil {
				fmt.Println(err.Error())
				delete(*conns, key)
			}
			checkError(err, "数据写入信息：")

		}
	}

}

////////////////////////////////////////////////////////
//
//启动服务器
//参数
//  端口 port
//
////////////////////////////////////////////////////////
func StartServer(port string) {
	service := ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err, "接收tcp信息:")
	l, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "监听tcp信息:")
	conns := make(map[string]net.Conn)
	messages := make(chan string, 10)

	//启动服务器广播线程
	go FwqSend(&conns, messages)

	for {
		fmt.Println("开始监听：")
		conn, err := l.Accept()
		checkError(err, "开始接收：")
		fmt.Println("正在接收中。。。")
		conns[conn.RemoteAddr().String()] = conn
		//启动一个新线程
		go FwqReceive(conn, messages)

	}

}
