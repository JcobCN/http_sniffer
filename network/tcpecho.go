package network

import (
	"bufio"
	"fmt"
	"net"
)

func TcpServer(){
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:32333")
	if err != nil{
		fmt.Printf("err %+v\n",err)
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil{
		fmt.Printf("err %+v\n", err)
	}
	defer tcpListener.Close()

	fmt.Println("等待客户端连接....")

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil{
			fmt.Printf("接受客户端出错%+v\n", err)
			continue
		}

		fmt.Println("一个客户端已连接："+tcpConn.RemoteAddr().String())

		//启用一个gorutine
		go tcpHandle(tcpConn)
	}
}

func tcpHandle(conn *net.TCPConn){
	ipStr := conn.RemoteAddr().String()

	defer func(){
		fmt.Println("断开连接："+ipStr)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	//i := 0

	//接收消息
	for{
		message, err := reader.ReadByte()
		if err != nil{
			fmt.Println("读取信息出错")
			break
		}
		fmt.Println(string(message))
	}

}

func TcpClient(addr string){
	conn, err := net.Dial("tcp", addr+)
	if err != nil{
		fmt.Println(addr+"连接失败：", err)
		return
	}
	defer conn.Close()

}
