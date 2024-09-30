package main

import (
	"fmt"
	"time"

	"github.com/chinese-room-solutions/fakebro/useragent"
)

func main() {
	seed := time.Now().UnixNano()
	fmt.Printf("Seed: %d\n", seed)
	headers := useragent.NewUserAgent(20, seed).Headers

	fmt.Println("Generated Headers:")
	for key, value := range headers {
		fmt.Printf("%s: %s\n", key, value)
	}
}
