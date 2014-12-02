package main

import (
	"fmt"

	"github.com/godbus/dbus"
)

type Server struct {
	conn *dbus.Conn
}

func NewServer() (*Server, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}
	// XXX: reply処理
	_, err = conn.RequestName("org.freedesktop.Notifications", dbus.NameFlagDoNotQueue)
	if err != nil {
		return nil, err
	}

	s := &Server{conn}

	conn.Export(s, "/org/freedesktop/Notifications", "org.freedesktop.Notifications")

	return s, nil
}

// action-icons, actions, body, body-hyperlinks, body-images, body-markup, icon-multi, icon-static, persistence, sound
func (s Server) GetCapabilities() ([]string, *dbus.Error) {
	// TODO: select return variables
	return []string{"body"}, nil
}

func (s Server) Notify(appName string, replacesID uint32, appIcon, summary, body string, actions []string, hints map[string]dbus.Variant, expireTimeout int32) (uint32, *dbus.Error) {
	// TODO: result
	// TODO: Do notify
	fmt.Println(summary)
	fmt.Println(body)
	return 1, nil
}

func (s Server) CloseNotification(id uint32) *dbus.Error {
	// TODO: delete Notification
	//       Emitte NotificationClosed signal.
	//       if notification dosen't exists, return empty dbus-error
	return nil
}

func (s Server) GetServerInformation() (name, vendor, version, specVersion string, err *dbus.Error) {
	// TODO: Hard coded result
	return "GoNotify", "pocke", "0.0.1", "1.2", nil
}

func main() {
	_, err := NewServer()
	if err != nil {
		panic(err)
	}

	fmt.Println("start")

	select {}
}
