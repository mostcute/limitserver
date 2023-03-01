package main

import (
	"fmt"
	"github.com/mostcute/limitserver/client"
	"log"
)

func main() {
	fmt.Println("start")
	c := client.NewKmClient([]string{"127.0.0.1:7777"})
	fmt.Println("end")
	for {
		err := c.GetToken()
		if err != nil {
			log.Println("err", err)
		}
	}

}
