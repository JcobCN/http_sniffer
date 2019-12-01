package tcp_handle

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

var homePath = os.Getenv("HOMEDRIVE")+os.Getenv("HOMEPATH")



func RemoteHandle(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()

	defer func() {
		fmt.Println("断开连接：" + ipStr)
		conn.Close()
	}()


	message := make([]byte, 4096)

	//接收消息
	for {
		n, err := conn.Read(message)
		if err != nil{
			fmt.Println(err)
		}

		fmt.Print("--->:" + string(message[:n]))
		caseString := string(message[:n])
		switch caseString{
		case "stop_write_file":

		}


	}

}

func getMsgFromCLI(conn *net.TCPConn) error {

	stdi := bufio.NewReader(os.Stdin)

	fmt.Print("请输入：")
	wmsg, _ := stdi.ReadString('\n')
	fmt.Print("<---:" + wmsg)

	_, err := conn.Write([]byte(wmsg))
	return err
}

func ControllerHandle(conn *net.TCPConn) {
	b := []byte("Hello control server\n")
	conn.Write(b)

	reader := bufio.NewReader(conn)
	for {
		rmsg, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			fmt.Println("读取信息出错", rmsg, err)
			break
		}
		fmt.Print("--->:" + string(rmsg))

		err = getMsgFromCLI(conn)
		if err != nil {
			fmt.Println(err)
			break
		}

	}
}
