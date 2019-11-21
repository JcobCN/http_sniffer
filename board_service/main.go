package main

import (
	"fmt"
	"httpsniffer/network"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Pls, input argument which is 'server' or 'connect'")
		return
	}

	switch os.Args[1] {
	case "server":
		network.BoardCastServer()

		//client
	case "connect":
		network.BoardCastClient()
	default:
		fmt.Println(os.Args[0] + " arguments is 'server' or 'connect'")
	}
}
