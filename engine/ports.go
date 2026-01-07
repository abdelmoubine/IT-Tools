package engine

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type PortResult struct {
	IP     string
	Port   int
	State  string
	Banner string
}

// PortScan scans ports for a single IP and port spec (e.g. "1-1024" or "22,80,443")
func PortScan(ip string, portSpec string) []PortResult {
	ports := parsePortSpec(portSpec)
	timeout := 300 * time.Millisecond
	var wg sync.WaitGroup
	results := make([]PortResult, 0)
	resLock := sync.Mutex{}

	for _, p := range ports {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := net.JoinHostPort(ip, strconv.Itoa(port))
			conn, err := net.DialTimeout("tcp", address, timeout)
			if err != nil {
				// closed/filtered
				resLock.Lock()
				results = append(results, PortResult{IP: ip, Port: port, State: "closed"})
				resLock.Unlock()
				return
			}
			conn.Close()
			resLock.Lock()
			results = append(results, PortResult{IP: ip, Port: port, State: "open"})
			resLock.Unlock()
		}(p)
	}
	wg.Wait()
	return results
}

func parsePortSpec(spec string) []int {
	spec = strings.TrimSpace(spec)
	if strings.Contains(spec, "-") {
		parts := strings.Split(spec, "-")
		if len(parts) == 2 {
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			out := make([]int, 0)
			for i := start; i <= end; i++ {
				out = append(out, i)
			}
			return out
		}
	}
	// comma separated
	parts := strings.Split(spec, ",")
	out := make([]int, 0)
	for _, p := range parts {
		v, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil {
			out = append(out, v)
		}
	}
	return out
}