package engine

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"
)

// QuickScan accepts a CIDR or comma list; returns slice of reachable IPs (by TCP connect to common port).
func QuickScan(target string) []string {
	// Very simple: if CIDR, expand; if single IP or comma list, iterate.
	targets := parseTargets(target)
	results := make([]string, 0)
	timeout := 500 * time.Millisecond
	for _, t := range targets {
		// try common port 80 then 443
		ok := false
		for _, p := range []string{"80", "443", "22"} {
			if tcpConnect(t, p, timeout) {
				ok = true
				break
			}
		}
		if ok {
			results = append(results, t)
		}
	}
	return results
}

func parseTargets(target string) []string {
	target = strings.TrimSpace(target)
	if strings.Contains(target, "/") {
		ips, _ := hostsFromCIDR(target)
		return ips
	}
	if strings.Contains(target, ",") {
		parts := strings.Split(target, ",")
		out := make([]string, 0)
		for _, p := range parts {
			out = append(out, strings.TrimSpace(p))
		}
		return out
	}
	return []string{target}
}

func tcpConnect(host, port string, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	d := net.Dialer{}
	conn, err := d.DialContext(ctx, "tcp", net.JoinHostPort(host, port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// expand CIDR to hosts (small networks only)
func hostsFromCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		ips = append(ips, ip.String())
	}
	// remove network & broadcast
	if len(ips) > 2 {
		return ips[1 : len(ips)-1], nil
	}
	return ips, nil
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func Traceroute(host string, maxHops int) ([]string, error) {
	// Simple stub: try system "tracert" on Windows
	out := make([]string, 0)
	// We'll use net.LookupIP as placeholder
	ips, err := net.LookupIP(host)
	if err == nil {
		for _, ip := range ips {
			out = append(out, ip.String())
		}
		return out, nil
	}
	return []string{"traceroute not available"}, err
}