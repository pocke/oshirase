package main

import (
	"fmt"

	"github.com/pocke/oshirase"
)

var f = func(n *oshirase.Notify) {
	fmt.Println(n.ID)
	fmt.Println(n.Summary)
	fmt.Println(n.Body)
}

func main() {
	srv, err := oshirase.NewServer("Notification Server", "Pocke", "0.0.1")
	if err != nil {
		panic(err)
	}

	srv.OnNotify(f)
	srv.OnCloseNotification(func(_ uint32) bool { return true })

	fmt.Println("start")

	select {}
}
