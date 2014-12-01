package main

import (
	"fmt"

	"github.com/godbus/dbus"
)

type server string

// action-icons, actions, body, body-hyperlinks, body-images, body-markup, icon-multi, icon-static, persistence, sound
func (s server) GetCapabilities() ([]string, *dbus.Error) {
	return []string{"body"}, nil
}

func (s server) Notify(appName string, replacesID uint32, appIcon, summary, body string, actions []string, hints map[string]dbus.Variant, expireTimeout int32) (uint32, *dbus.Error) {
	fmt.Println(summary)
	fmt.Println(body)
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
