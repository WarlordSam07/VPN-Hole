package main

import (
	"fmt"
	"log"

	"0xacab.org/leap/vpn-hole/vpnhole/vpnhole"
	"github.com/common-nighthawk/go-figure"
)

func main() {

	// Print the 'VPN-Hole' Logo
	Logo := figure.NewColorFigure(" VPN-Hole", "", "green", true)
	Logo.Print()

	fmt.Println("\033[32m[+] \033[0m===============================================================\033[32m[+]\033[0m")

	c := vpnhole.NewVpnHoleClient("", "", "", nil)

	// start the vpnhole
	_, err := c.Start()
	if err != nil {
		log.Fatal(err)
	}

	// stop the vpnhole
	defer c.Stop()
	log.Println("\033[31m[+] Stopping VPN-Hole\033[0m")
}
