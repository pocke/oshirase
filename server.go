package oshirase

import (
	"fmt"

	"github.com/godbus/dbus"
)

type Server struct {
	conn                *dbus.Conn
	name                string
	vendor              string
	version             string
	onNotify            func(*Notify)
	onCloseNotification func(uint32) bool
	notifyID            <-chan uint32
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

	notifyID := make(chan uint32)
	go func(ch chan uint32) {
		var i uint32
		for i = 1; ; i++ {
			ch <- i
		}
	}(notifyID)

	s := &Server{
		conn:     conn,
		name:     name,
		vendor:   vendor,
		version:  version,
		notifyID: (<-chan uint32)(notifyID),
	}

	m := &messages{server: s}

	conn.Export(m, "/org/freedesktop/Notifications", "org.freedesktop.Notifications")

	return s, nil
}

func (s *Server) Close() error {
	return s.conn.Close()
}

func (s *Server) OnNotify(f func(*Notify)) {
	s.onNotify = f
}

func (s *Server) OnCloseNotification(f func(uint32) bool) {
	s.onCloseNotification = f
}
