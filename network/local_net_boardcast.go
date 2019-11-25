package network

import (
	"fmt"
	"net"
	"time"
)

func BoardCastServer()  {
	//local 缺省
	// 这里设置发送者的IP地址，自己查看一下自己的IP自行设定
	//laddr := net.UDPAddr{
	//	//IP:   net.IPv4(192, 168, 137, 224),
	//	//IP:   net.IPv4(0, 0, 0, 0),
	//	IP: net.IPv4zero,
	//	Port: 3000,
	//}

	//remote
	// 这里设置接收者的IP地址为广播地址
	raddr := net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 3000,
	}

	timeout := time.NewTimer(time.Second * 10)

	for{
		select {
		case <-timeout.C:
			go func() {
				conn, err := net.DialUDP("udp", nil, &raddr)
				if err != nil {
					fmt.Println("服务未开启", err)
					return
				}
				defer conn.Close()


				_, err = conn.Write([]byte("hello peers\n"))
				if err != nil{
					fmt.Println("发送失败:", err)
				}
			}()
			timeout.Reset(time.Second*10)
		}
	}

}

func BoardCastClient() error {
	addr,err := net.ResolveUDPAddr("udp","0.0.0.0:3000")
	if err != nil{
		fmt.Println("解析地址失败")
		return err
	}
	fmt.Println("解析地址：",addr.IP.String(),":",addr.Port)

	// 循环监听广播包，暂时不需要
	//for {
		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			fmt.Println("创建监听conn失败")
			return err
		}

		data := make([]byte, 1024)
		n, laddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data[:n]), laddr, n)
		conn.Close()
	//}



	return err
}