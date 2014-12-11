package oshirase

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
