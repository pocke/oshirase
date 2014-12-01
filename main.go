package main

import (
	"fmt"

	"github.com/godbus/dbus"
)

type server string

func (s server) Notify(app_name string, replaces_id uint32, app_icon, summary, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) (uint32, *dbus.Error) {
	fmt.Println(summary)
	return 1, nil
}

func (s server) GetServerInformation() (name, vendor, version, specVersion string, err *dbus.Error) {
	return "GoNotify", "pocke", "0.0.1", "1.2", nil
}

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	// XXX: reply処理
	_, err = conn.RequestName("org.freedesktop.Notifications", dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}

	n := server("Notify!!")
	conn.Export(n, "/org/freedesktop/Notifications", "org.freedesktop.Notifications")

	fmt.Println("start")

	select {}
}
