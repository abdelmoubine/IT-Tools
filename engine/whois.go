package engine

import (
	"net"
	"strings"
	"time"
)

// Basic whois lookup (very simple)
func WhoisLookup(query string) (string, error) {
	// Resolve to IP if hostname
	ips, err := net.LookupIP(query)
	if err == nil && len(ips) > 0 {
		query = ips[0].String()
	}
	// Use whois.iana.org as basic server lookup chain — this is minimal and may not return full data
	resp, err := queryWhois("whois.iana.org", query)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(resp), nil
}

func queryWhois(server, q string) (string, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(server, "43"), 5*time.Second)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	_, err = conn.Write([]byte(q + "\r\n"))
	if err != nil {
		return "", err
	}
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	return string(buf[:n]), nil
}