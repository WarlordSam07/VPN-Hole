package vpnhole

import (
	"errors"
	"fmt"
	"log"
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
		InfoLogger:  log.New(log.Writer(), "INFO: ", log.Flags()),
		ErrorLogger: log.New(log.Writer(), "ERROR: ", log.Flags()),
	}
}

// NewVpnHoleClient creates a new VpnHoleClient

func NewVpnHoleClient(addr, subscriptionsFilename, upstream string, eventlogger EventLogger) *VpnHoleClient {
	return &VpnHoleClient{
		Addr:                  addr,
		SubscriptionsFilename: subscriptionsFilename,
		Upstream:              upstream,
		Eventlogger:           eventlogger,
	}
}

// Start starts the VpnHoleClient and config new values
func (c *VpnHoleClient) Start() error {

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
		c.Eventlogger = &DefaultEventLogger{
			InfoLogger:  log.New(log.Writer(), "INFO: ", log.Flags()),
			ErrorLogger: log.New(log.Writer(), "ERROR: ", log.Flags()),
		}
	}

	c.Eventlogger.Info("Starting VPN-Hole")

	if c.IsRunning() {
		return ErrAlreadyRunning
	}

	return nil
}

// Stop stops the VpnHoleClient
func (c *VpnHoleClient) Stop() error {
	if !c.IsRunning() {
		return ErrNotRunning
	}

	c.Eventlogger.Error("Stopping VPN-Hole")

	return nil
}

// IsRunning returns true if the VpnHoleClient is running
func (c *VpnHoleClient) IsRunning() bool {
	return false
}
