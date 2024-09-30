package main

import (
	"fmt"
	"time"

	"github.com/chinese-room-solutions/fakebro/useragent"
)

func main() {
	// seed := time.Now().UnixNano()
	// fmt.Printf("Seed: %d\n", seed)
	// headers := useragent.NewUserAgent(20, seed, func(tt useragent.TokenType) bool {
	// 	return true
	// }).Headers

	// fmt.Println("Generated Headers:")
	// for key, value := range headers {
	// 	fmt.Printf("%s: %s\n", key, value)
	// }

	// Generate a 1000 user agents and measure the time it takes
	start := time.Now()
	for i := 0; i < 1000; i++ {
		seed := time.Now().UnixNano()
		useragent.NewUserAgent(20, seed, func(tt useragent.TokenType) bool {
			return true
		})
	}
	fmt.Printf("Generated 1000 user agents in %s\n", time.Since(start))
}
