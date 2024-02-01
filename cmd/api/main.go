package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("API running...")
		time.Sleep(10 * time.Second)
	}
}
