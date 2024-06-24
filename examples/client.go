package main

import (
	"log"
	"os"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {

	m := metrics.NewMeter()
	metrics.Register("quux", m)
	m.Mark(1)

	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	var j int64
	j = 5
	for true {
		time.Sleep(time.Second * 1)

		m.Mark(j)
	}
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/mostcute/limitserver/client"
// )

// func main() {
// 	fmt.Println("start")
// 	c := client.NewLimitClient([]string{"127.0.0.1:7777"})
// 	fmt.Println("end")
// 	// fmt.Println("Total == ", c.Limit())
// 	i := 0
// 	lasti := 0
// 	go func() {
// 		for {
// 			time.Sleep(time.Second)
// 			fmt.Println("last call", i-lasti)
// 			lasti = i
// 			usg, err := c.Usage()
// 			if err != nil {
// 				log.Println("err", err)
// 			}
// 			fmt.Println("usg", usg)
// 		}
// 	}()
// 	for {
// 		_, err := c.GetToken()
// 		i++
// 		if err != nil {
// 			log.Println("err", err)
// 		}
// 		// fmt.Println("token", token)
// 	}

// }
