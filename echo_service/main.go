package main

import (
	"fmt"
	"httpsniffer/network"
	"os"
)

func main(){
	switch os.Args[1] {
	case "server":
		network.TcpServer()
	case "client":
		if os.Args[2] == ""{
			fmt.Println("please input remote ip addr, behind the 'client' command. ")
			return
		}
		network.TcpClient(os.Args[2])
	default:
		fmt.Println(os.Args[0] + " arguments is 'server' or 'client'")
	}
}
