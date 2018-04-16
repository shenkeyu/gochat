package main

import (
	"code/skygo/firstgo"
	"code/skygo/fuwuqi"
	"code/skygo/kehuduan"
	"fmt"
	"os"
)

////////////////////////////////////////////////////////
//
//主程序
//
//参数说明：
//  启动服务器端：  Chat server [port]             eg: Chat server 9090
//  启动客户端：    Chat client [Server Ip Addr]:[Server Port]    eg: Chat client 192.168.0.74:9090
//
////////////////////////////////////////////////////////
func main() {
	firstgo.Tryfirst()
	fmt.Printf("**************************************\n")
	fmt.Printf("        欢迎来到王者荣耀聊天室            \n")
	fmt.Printf("**************************************\n")
	fmt.Printf("（请注意文明聊天。有问题请输入‘/help’进行解答）\n")
	if len(os.Args) != 3 {
		fmt.Println("错误的输入！")
		os.Exit(0)
	}

	if os.Args[1] == "server" && len(os.Args) == 3 {

		fuwuqi.StartServer(os.Args[2])
	}

	if os.Args[1] == "client" && len(os.Args) == 3 {

		kehuduan.StartClient(os.Args[2])
	}

}
