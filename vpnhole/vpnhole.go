package vpnhole

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/miekg/dns"
)

var (
	ErrAlreadyRunning = errors.New("already running")
	ErrNotRunning     = errors.New("not running")
)

type VpnHoleClient struct {
	Addr                  string
	SubscriptionsFilename string
	Upstream              string
	Eventlogger           EventLogger
	server                *dns.Server
}

func (c *VpnHoleClient) String() string {
	return fmt.Sprintf("vpnhole client: addr=%s, upstream=%s, subscriptions=%s", c.Addr, c.Upstream, c.SubscriptionsFilename)
}

// EventLogger is an interface for logging events.
type EventLogger interface {
	Info(msg string)
	Error(msg string)
}

// DefaultEventLogger is the default event logger.
type DefaultEventLogger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

// Info logs an info message.
func (l *DefaultEventLogger) Info(msg string) {
	l.InfoLogger.Println(msg)
}

// Error logs an error message.
func (l *DefaultEventLogger) Error(msg string) {
	l.ErrorLogger.Println(msg)
}

// NewDefaultEventLogger returns a new DefaultEventLogger.
func NewDefaultEventLogger() *DefaultEventLogger {
	return &DefaultEventLogger{
		InfoLogger:  log.New(log.Writer(), "\033[32m[+] ", log.Flags()),
		ErrorLogger: log.New(log.Writer(), "\033[31m[-] ", log.Flags()),
	}
}

// NewVpnHoleClient creates a new VpnHoleClient

func NewVpnHoleClient(addr, subscriptionsFilename, upstream string, eventlogger EventLogger) *VpnHoleClient {
	return &VpnHoleClient{
		Addr:                  addr,
		SubscriptionsFilename: subscriptionsFilename,
		Upstream:              upstream,
		Eventlogger:           eventlogger,
		server:                &dns.Server{},
	}
}

// Start starts the VpnHoleClient and config new values
func (c *VpnHoleClient) Start() (bool, error) {

	defer func() {
		log.Println("\033[31m[+] Stopping VPN-Hole\033[0m")
	}()

	if c.IsRunning() {
		log.Println("[-] VPN-Hole is already running")
		return false, ErrAlreadyRunning
	}

	log.Println("\033[32m[+] Starting VPN-Hole\033[0m")

	if c.Addr == "" {
		c.Addr = ":53"
	}

	if c.SubscriptionsFilename == "" {
		c.SubscriptionsFilename = "subs.list"
	}

	if c.Upstream == "" {
		c.Upstream = "1.1.1.1:53"
	}

	if c.Eventlogger == nil {
		c.Eventlogger = NewDefaultEventLogger()
	}

	c.server = &dns.Server{Addr: c.Addr}

	subscriptions, err := ReadSubscriptions(c.SubscriptionsFilename)
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to read subscriptions list: %w", err))

		return false, err
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

		c.Addr = ""
		return false, err
	}

	return true, nil
}

// Stop stops the VpnHoleClient
func (c *VpnHoleClient) Stop() (bool, error) {

	if !c.IsRunning() {
		return false, ErrNotRunning
	}

	if err := c.server.Shutdown(); err != nil {
		log.Fatalln(fmt.Errorf("failed to shutdown DNS server: %w", err))
		return false, err
	}

	c.server = nil
	c.Addr = ""

	return true, nil
}

// IsRunning returns true if the VpnHoleClient is running.
func (c *VpnHoleClient) IsRunning() bool {
	return c.Addr != "" && c.server != nil
}
