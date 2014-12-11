package main

import (
	"fmt"

	"github.com/godbus/dbus"
)

type Server struct {
	conn    *dbus.Conn
	name    string
	vendor  string
	version string
}

const (
	CloseReasonExpired uint32 = iota + 1
	CloseReasonDismissedByUser
	CloseReasonCallToCloseNotification
	CloseReasonUndefinedOrReserved
)

const (
	specVersion string = "1.2"
)

func NewServer(name, vendor, version string) (*Server, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}
	reply, err := conn.RequestName("org.freedesktop.Notifications", dbus.NameFlagDoNotQueue)
	if err != nil {
		return nil, err
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		return nil, fmt.Errorf("name already token")
	}

	s := &Server{
		conn:    conn,
		name:    name,
		vendor:  vendor,
		version: version,
	}

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
	// TODO: Do delete Notification, and return delete success?
	//       if notification dosen't exists, return empty dbus-error
	s.conn.Emit("/org/freedesktop/Notifications/NotificationClosed", "org.freedesktop.Notifications.NotificationClosed", id, CloseReasonCallToCloseNotification)
	return nil
}

func (s Server) GetServerInformation() (name, vendor, version, specVer string, err *dbus.Error) {
	return s.name, s.vendor, s.version, specVersion, nil
}

func main() {
	_, err := NewServer("pocke", "pocket", "0.0.1")
	if err != nil {
		panic(err)
	}

	fmt.Println("start")

	select {}
}
