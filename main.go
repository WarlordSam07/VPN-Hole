package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/miekg/dns"
)

var httpClient = &http.Client{Timeout: 30 * time.Second}

func main() {
	var (
		addr                  string
		subscriptionsFilename string
	)

	flag.StringVar(&addr, "addr", ":5353", "addr")
	flag.StringVar(&upstream, "upstream", "1.1.1.1:5353", "upstream")
	flag.StringVar(&subscriptionsFilename, "subs", "subs.list", "subscriptions filename")
	flag.Parse()

	subscriptions, err := readSubscriptions(subscriptionsFilename)
	if err != nil {
		log.Fatalln(fmt.Errorf("Failed to read subscriptions list: %w", err))
	}

	for _, blacklistURL := range subscriptions {
		privBlacklist.Subscribe(blacklistURL)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go privBlacklist.Watch(ctx, time.Minute*10)

	dns.HandleFunc(".", handler)

	if err = dns.ListenAndServe(addr, "udp", nil); err != nil {
		log.Println(fmt.Errorf("Failed to serve DNS server: %w", err))
	}
}
