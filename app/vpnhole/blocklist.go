package main

import (
	"net/http"
	"strings"
	"time"

	"0xacab.org/leap/vpn-hole/vpnhole/blacklist"
	"github.com/miekg/dns"
)

var httpClient = &http.Client{Timeout: 30 * time.Second}
var PrivBlacklist = blacklist.New(httpClient)

func IsBlacklisted(req *dns.Msg) bool {
	if req.Opcode != dns.OpcodeQuery {
		return false
	}

	if len(req.Question) != 1 {
		return false
	}

	q := req.Question[0]

	switch q.Qtype {
	case dns.TypeA:
	case dns.TypeAAAA:
	default:
		return false
	}

	return PrivBlacklist.Contains(strings.TrimRight(q.Name, "."))
}
