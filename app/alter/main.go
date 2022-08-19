package main

import (
	"context"
	"fmt"
	"log"

	"0xacab.org/leap/vpn-hole/alter/alter"
	"github.com/miekg/dns"

	"time"
)

func init() {
	// start the cmdline parser
	alter.ParseFlags()

}

func main() {

	// call ParseFlags() to get the config struct with the values
	c := alter.ParseFlags()
	fmt.Println(c)

	// start the vpnhole
	if err := c.Start(); err != nil {
		log.Fatalln(fmt.Errorf("failed to start vpnhole: %w", err))
	}
	defer func() {
		if err := c.Stop(); err != nil {
			log.Println(fmt.Errorf("failed to stop vpnhole: %w", err))
		}
	}()
	log.Printf("vpnhole started: %s", c)

	subscriptions, err := ReadSubscriptions(c.SubscriptionsFilename)
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to read subscriptions list: %w", err))
	}

	for _, blacklistURL := range subscriptions {
		PrivBlacklist.Subscribe(blacklistURL)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go PrivBlacklist.Watch(ctx, time.Minute*10)

	dns.HandleFunc(".", Handler)

	if err = dns.ListenAndServe(c.Addr, "udp", nil); err != nil {
		log.Println(fmt.Errorf("failed to serve DNS server: %w", err))
	}

}
