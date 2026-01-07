package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("IT Support Toolkit Agent (scaffold) starting...")
	for {
		// Poll management server for updates (stub)
		fmt.Println("Agent poll at", time.Now())
		time.Sleep(30 * time.Second)
	}
}