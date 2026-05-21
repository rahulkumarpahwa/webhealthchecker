package checker

import (
	"fmt"
	"net"
	"time"
)

type Check struct {
	Domain string `json:"domain"`
	Port   string `json:"port"`
}

func (c *Check) Checker() {
	addr := c.Domain + ":" + c.Port
	conn, err := net.DialTimeout(
		"tcp",
		addr,
		5*time.Second,
	)

	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}


	fmt.Println("Connected")

	fmt.Printf("Connection Successfull from: %v to %v!", conn.LocalAddr(), conn.RemoteAddr())
	defer conn.Close()
}
