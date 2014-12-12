package main

import (
	"fmt"

	"github.com/pocke/oshirase"
)

var f = func(n *oshirase.NotifyArg) error {
	fmt.Println(n.ID)
	fmt.Println(n.Summary)
	fmt.Println(n.Body)
	return nil
}

func main() {
	srv, err := oshirase.NewServer("Notification Server", "Pocke", "0.0.1")
	if err != nil {
		panic(err)
	}

	srv.OnNotify(f)

	fmt.Println("start")

	select {}
}
