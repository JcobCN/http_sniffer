package main

import (
	"fmt"
	"github.com/google/gopacket"
	_ "github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"httpsniffer/file"
	"log"
	"os"
	"strings"
	"time"
)

var DEBUG int = 0

func snifferHttp(){
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	var netCardName string

	// Print device information
	fmt.Println("Devices found:")
	for _, d := range devices {
		fmt.Println("\nName: ", d.Name)
		fmt.Println("Description: ", d.Description)
		fmt.Println("Devices addresses: ", d.Description)

		if strings.Contains( strings.ToLower(d.Description), "pcie") {
			netCardName = d.Name
			break
		}

		for _, address := range d.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}

	netCardName1 := "\\Device\\NPF_{A03165A0-9781-4C35-8298-FEC0E040754A}"
	netCardName2 := "\\Device\\NPF_{FFBAE5D2-88B7-4311-BE9A-09335FA9F87D}"

	switch DEBUG{
	case 0:
		netCardName = netCardName
	case 1:
		netCardName = netCardName1
	case 2:
		netCardName = netCardName2
	}

	handle, err := pcap.OpenLive(netCardName, 1600, true, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	//设置过滤
	if err := handle.SetBPFFilter("tcp and (port 80)"); err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	packets := packetSource.Packets()

	for {
		select{
		case packet := <-packets:
			if packet == nil {
				return
			}

			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				log.Println("Unusable packet")
				continue
			}
			tcp := packet.TransportLayer().(*layers.TCP)
			payload := string(tcp.BaseLayer.Payload)
			if strings.Contains(payload, "GET") || strings.Contains(payload, "POST") {
				log.Printf("payload:%v\n", payload)
				homePath := os.Getenv("HOMEPATH")
				file.WriteWithOs(homePath+"/1.txt", payload)
			}

		}
	}
}

func main(){

	//snifferHttp()
}

