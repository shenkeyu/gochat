package kehuduan

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

////////////////////////////////////////////////////////
//
//客户端发送线程
//参数
//      发送连接 conn
//
////////////////////////////////////////////////////////
func chatSend(conn net.Conn) {

	var input string
	username := conn.LocalAddr().String()
	for {

		fmt.Scanln(&input)
		if input == "/quit" {
			fmt.Println("聊天结束，再见！")
			conn.Close()
			os.Exit(0)
		}

		if input == "/help" {
			fmt.Println("/quit命令，为退出聊天！")
			fmt.Println("/search命令，为搜索你要的内容！")
			fmt.Println("/file命令，为退出聊天传送一个文件，图片等各类格式！")
		}
		str := []string{"/help", "/search", "/file"}
		boolstr := true
		for i, v := range str {
			if input == v {
				i++
				boolstr = false
			}
		}
		if boolstr {
			lens, err := conn.Write([]byte(username + " 说:" + input))
			fmt.Println("你已发送" + strconv.Itoa(lens) + "字节信息.")
			if err != nil {
				fmt.Println(err.Error())
				conn.Close()
				break
			}
		}

	}

}

////////////////////////////////////////////////////////
//
//客户端启动函数
//参数
//      远程ip地址和端口 tcpaddr
//
////////////////////////////////////////////////////////
func StartClient(tcpaddr string) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpaddr)
	checkError(err, "处理tcp地址信息：")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, "连接tcp信息：")
	//启动客户端发送线程
	go chatSend(conn)

	//开始客户端轮训
	buf := make([]byte, 1024)
	for {

		lenght, err := conn.Read(buf)
		if checkError(err, "连接信息：") == false {
			conn.Close()
			fmt.Println("服务器停止服务，再见！")
			os.Exit(0)
		}
		fmt.Println(string(buf[0:lenght]))

	}
}
