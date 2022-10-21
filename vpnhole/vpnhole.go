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

type EventLogger interface {
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type DefaultEventLogger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

// Implement the EventLogger interface
func (l *DefaultEventLogger) Infof(format string, v ...interface{}) {
	l.InfoLogger.Printf(format, v...)
}

func (l *DefaultEventLogger) Errorf(format string, v ...interface{}) {
	l.ErrorLogger.Printf(format, v...)
}

func (c *VpnHoleClient) String() string {
	return fmt.Sprintf("addr=%s, subscriptions=%s, upstream=%s", c.Addr, c.SubscriptionsFilename, c.Upstream)
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

	c.Eventlogger.Infof("vpnhole started: %s", c)

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

	return nil
}

// IsRunning returns true if the VpnHoleClient is running
func (c *VpnHoleClient) IsRunning() bool {
	return false
}
