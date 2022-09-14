package vpnhole

import (
	"fmt"
)

var ErrShutdown = fmt.Errorf("vpnhole was shutdown gracefully")

type vpnholeclient struct {
	Addr                  string
	SubscriptionsFilename string
	Upstream              string
}

// parse the flags and return the config struct with the values
func ParseFlags() vpnholeclient {
	return vpnholeclient{
		Addr:                  ":53",
		SubscriptionsFilename: "subs.list",
		Upstream:              "1.1.1.1:53",
	}
}

// start the vpnhole
func (c *vpnholeclient) Start() error {
	return nil
}

// stop the vpnhole
func (c *vpnholeclient) Stop() error {
	return nil
}
