package hardware

import (
	"fmt"
	"log"
	"net"
	"os"
)

func GetMacAddrs() (macAddrs string ) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = macAddrs + "," + macAddr
	}
	return macAddrs
}

func GetIPs() (ips string) {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = ips +","+ipNet.IP.String()
			}
		}
	}
	return ips
}


func GetComName() string {
	host, err := os.Hostname()
	if err != nil{
		fmt.Printf("%s\n", err)
	}else{
		fmt.Printf("%s\n", host)
	}

	return host
}


func GetMacAddr(){
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		fmt.Println(inter.Name)
		fmt.Println(inter.Flags)
		fmt.Println(inter)
		mac := inter.HardwareAddr //获取本机MAC地址
		fmt.Println("MAC = ", mac)
	}
}

func a(){
	//以太网网卡名称为eth0
	inter, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Fatalln(err)
	}
	//mac地址
	fmt.Println(inter.HardwareAddr.String())
	addrs, err := inter.Addrs()
	if err != nil {
		log.Fatalln(err)
	}
	//ip地址一个ip4一个ip6
	for _, addr := range addrs {
		fmt.Println(addr.String())
	}
}