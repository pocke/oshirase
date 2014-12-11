package main

import (
	"fmt"

	"github.com/pocke/oshirase"
)

func main() {
	_, err := oshirase.NewServer("Notification Server", "Pocke", "0.0.1")
	if err != nil {
		panic(err)
	}

	fmt.Println("start")

	select {}
}
