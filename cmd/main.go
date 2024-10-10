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

	renderer := webgl.NewWebGLRenderer("Windows", "10.0.0", 12345)
	fmt.Println(renderer.Renderer)
}
