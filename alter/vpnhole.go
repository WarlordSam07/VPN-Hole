package alter

import (
	"fmt"
)

var ErrShutdown = fmt.Errorf("vpnhole was shutdown gracefully")

type Config struct {
	Addr                  string
	SubscriptionsFilename string
	Upstream              string
	Start                 func() error
	Stop                  func() error
}

func (c Config) String() string {
	return fmt.Sprintf("Config{Addr: %s, SubscriptionsFilename: %s, Upstream: %s}", c.Addr, c.SubscriptionsFilename, c.Upstream)
}

// parse the flags and return the config struct with the values
func ParseFlags() Config {
	return Config{
		Addr:                  ":53",
		SubscriptionsFilename: "subs.list",
		Upstream:              "1.1.1.1:53",

		Start: func() error {
			return nil
		},
		Stop: func() error {
			return nil
		},
	}
}
