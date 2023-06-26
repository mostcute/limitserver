package main

import "github.com/mostcute/limitserver/server"

func main() {
	s := server.NewLimitService("7777")
	s.NewSpeedLimiter(3, 1, "aaa")
	s.Run()
}
