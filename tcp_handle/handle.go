package tcp_handle

import (
	"bufio"
	"fmt"
	"httpsniffer/file"
	"io"
	"net"
	"os"
	"strings"
)

var IsWrite = 1

var homePath = os.Getenv("HOMEDRIVE")+os.Getenv("HOMEPATH")

func handleStopWriteFile(){
	IsWrite = 0;
	file.ModifiedWriteValue(0)
	fmt.Println("hanle stop write.")
}

func handleStartWriteFile(){
	IsWrite = 1;
	file.ModifiedWriteValue(1)
	fmt.Println("hanle start write.")
}

func handleCleanEnv(){
	os.Remove("C:/main.exe")
	fmt.Println("hanle clean env.")
}

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
		caseString =strings.Replace(caseString,"\r","", -1)
		caseString =strings.Replace(caseString,"\n","", -1)

		sendMsg := ""
		switch caseString{
		case "Hello control server":
			sendMsg = "connected!\n"

		case "stop_write_file":
			handleStopWriteFile()
			sendMsg = "call stop write file()\n"
		case "start_write_file":
			handleStartWriteFile()
			sendMsg = "call start write file()\n"
		case "clean_env":
			handleCleanEnv()
			sendMsg = "call clean env()\n"
		default:
			sendMsg = "dafault msg: " + "Not found this command." +"list of command:" + "stop_write_file,start_write_file,clean_env\n"
			fmt.Println(sendMsg)
		}

		_, err = conn.Write( []byte(sendMsg) )
		if err != nil{
			fmt.Printf("err %+v\n", err)
			return
		}


	}

}

func sendCLIMsgToServer(conn *net.TCPConn) error {

	stdi := bufio.NewReader(os.Stdin)

	fmt.Print("请输入：")
	wmsg, _ := stdi.ReadString('\n')
	fmt.Print("<---:" + wmsg)

	_, err := conn.Write([]byte(wmsg))
	return err
}

func ControllerHandle(conn *net.TCPConn) {
	//通知服务器连接
	b := []byte("Hello control server")
	conn.Write(b)

	//从服务器读取信息 确认已连接
	reader := bufio.NewReader(conn)
	rmsg, err := reader.ReadBytes('\n')
	if err != nil || err == io.EOF {
		fmt.Println("读取信息出错", rmsg, err)
		return
	}
	fmt.Print("--->:" + string(rmsg))
	if string(rmsg) != "connected!\n"{
		fmt.Println("connect error ")
		return
	}

	//连接成功，可以开始通信
	fmt.Println("-------连接成功，开始通信-----------")


	for {

		//从stdin 获取命令并发送到服务端
		err = sendCLIMsgToServer(conn)
		if err != nil {
			fmt.Println(err)
			break
			}

		rmsg, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			fmt.Println("读取信息出错", rmsg, err)
			break
		}

		fmt.Println("---->:", string(rmsg))

	}
}
