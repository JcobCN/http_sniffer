/*
echo_server，分离了handle和连接
 */
package network

import (
	"fmt"
	"net"
)

//TCPHandleFunc is a func pointer
type TCPHandleFunc func(*net.TCPConn)

func TcpRemote(handle TCPHandleFunc) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:32333")
	if err != nil {
		fmt.Printf("err %+v\n", err)
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Printf("err %+v\n", err)
	}
	defer tcpListener.Close()

	fmt.Println("等待客户端连接....")

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Printf("接受客户端出错%+v\n", err)
			continue
		}

		fmt.Println("一个客户端已连接：" + tcpConn.RemoteAddr().String())

		//启用一个gorutine
		go handle(tcpConn)
	}
}

func TcpController(addr string, handle TCPHandleFunc) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr+":32333")
	if err != nil {
		fmt.Printf("err %+v\n", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(addr+"连接失败：", err)
		return
	}
	defer conn.Close()

	fmt.Println(addr + " 连接成功,等待接收服务端信息....")

	handle(conn)
}
