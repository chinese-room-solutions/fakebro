package main

import (
	"fmt"
	"time"

	"github.com/chinese-room-solutions/fakebro/useragent"
)

func main() {
	seed := time.Now().UnixNano()
	fmt.Printf("Seed: %d\n", seed)
	headers := useragent.NewUserAgent(15, seed).Headers

	fmt.Println("Generated Headers:")
	for key, value := range headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	// start := time.Now()
	// for i := 0; i < 1000; i++ {
	// 	seed := time.Now().UnixNano()
	// 	useragent.NewUserAgent(15, seed)
	// }
	// fmt.Printf("Generated 1000 user agents in %s\n", time.Since(start))
}
