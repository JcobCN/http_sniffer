package network

import "net"

func boardCast()  {
	//local
	// 这里设置发送者的IP地址，自己查看一下自己的IP自行设定
	laddr := net.UDPAddr{
		//IP:   net.IPv4(192, 168, 137, 224),
		//IP:   net.IPv4(0, 0, 0, 0),
		IP: net.IPv4zero,
		Port: 3000,
	}
	//remote
	// 这里设置接收者的IP地址为广播地址
	raddr := net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 3000,
	}
	conn, err := net.DialUDP("udp", &laddr, &raddr)
	if err != nil {
		println(err.Error())
		return
	}
	conn.Write([]byte(`hello peers`))
	conn.Close()
}

func client(){

}