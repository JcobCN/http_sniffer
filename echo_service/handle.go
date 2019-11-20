package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func remoteHandle(conn *net.TCPConn) {
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

func controllerHandle(conn *net.TCPConn) {
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
