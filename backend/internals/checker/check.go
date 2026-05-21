package checker

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Check struct {
	Domain string `json:"domain"`
	Port   string `json:"port"`
}

const Instances int = 2

var checkQueue = make(chan *Check, 100)

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
	fmt.Printf("Connection Successfull from: %s :  %v to  %s : %v!", "", conn.LocalAddr(), c.Domain, conn.RemoteAddr())
	defer conn.Close()
}

func (c *Check) Runner() {
	var wg sync.WaitGroup
	for i := 0; i < Instances; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for check := range checkQueue {
		check.Checker()
	}
}

func (c *Check) AddCheck(domain string, port string) {
	check := &Check{
		Domain: domain,
		Port:   port,
	}

	select {
	case checkQueue <- check:
		fmt.Println("New Check Added in Queue!")
	default:
		fmt.Println("Failed to add the check!")
	}
}

func CloseCheck()  {
	close(checkQueue)
}