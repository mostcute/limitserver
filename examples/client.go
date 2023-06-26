package main

import (
	"fmt"
	"github.com/mostcute/limitserver/client"
	"log"
	"time"
)

func main() {
	fmt.Println("start")
	c := client.NewLimitClient([]string{"127.0.0.1:7777"})
	fmt.Println("end")
	fmt.Println("Total == ", c.Limit())
	i := 0
	lasti := 0
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("last call", i-lasti)
			lasti = i
		}
	}()
	for {
		err := c.GetToken()
		i++
		if err != nil {
			log.Println("err", err)
		}
	}

}
