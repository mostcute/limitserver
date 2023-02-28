package main

import "github/mostcute/limitserver/server"

func main() {
	s := server.NewLimitService("7777")
	s.NewSpeedLimiter(3, 1)
	s.Run()
}
