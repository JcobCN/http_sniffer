package network

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func FileServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:32333")
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
		go tcpServerHandle(tcpConn)
	}
}

func fileServerHandle(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()

	defer func() {
		fmt.Println("断开连接：" + ipStr)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	//i := 0

	stdi := bufio.NewReader(os.Stdin)

	//接收消息
	for {

		message, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			fmt.Println("读取信息出错", err)
			break
		}

		fmt.Print("--->:" + string(message))

		fmt.Print("请输入：")
		wmsg, _ := stdi.ReadString('\n')
		fmt.Print("<---:" + wmsg)

		_, err = conn.Write([]byte(wmsg))
		if err != nil {
			fmt.Println(err)
			break
		}
	}

}

func FileClient(addr string) {
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

	tcpClientHandle(conn)
}

func fileClientHandle(conn *net.TCPConn) {
	b := []byte(conn.LocalAddr().String() + "--->:Hello control server\n")
	conn.Write(b)

	reader := bufio.NewReader(conn)
	stdi := bufio.NewReader(os.Stdin)
	for {
		rmsg, err := reader.ReadBytes('\n')
		fmt.Print("--->:" + string(rmsg))

		if err != nil || err == io.EOF {
			fmt.Println("读取信息出错", rmsg, err)
			break
		}

		fmt.Print("请输入：")
		wmsg, _ := stdi.ReadString('\n')
		fmt.Print("<---:" + wmsg)

		_, err = conn.Write([]byte(wmsg))
		if err != nil {
			fmt.Println(err)
			break
		}

	}
}
