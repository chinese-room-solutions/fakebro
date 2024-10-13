package main

import (
	"fmt"
	"time"

	"github.com/chinese-room-solutions/fakebro/useragent"
	"github.com/chinese-room-solutions/fakebro/webgl"
)

func main() {
	seed := time.Now().UnixNano()
	fmt.Printf("seed: %d\n", seed)
	headers := useragent.NewUserAgent(15, seed).Headers

	for key, value := range headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	renderer, err := webgl.GenerateRenderer(seed, "Linux", "6.10.5")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(renderer)
}
