package main

import (
	"fmt"
	"httpsniffer/network"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Pls, input argument which is 'server' or 'connect [ip]'")
		return
	}

	switch os.Args[1] {
	case "server":
		// network.TcpServer()
		network.TcpRemote(remoteHandle)

		//client
	case "connect":
		if len(os.Args) < 3 {
			fmt.Println("please input remote ip addr, behind the 'connect' command. ")
			return
		}
		// network.TcpClient(os.Args[2])
		network.TcpController(os.Args[2], controllerHandle)
	default:
		fmt.Println(os.Args[0] + " arguments is 'server' or 'connect'")
	}
}
