package oshirase

import "github.com/godbus/dbus"

type Notify struct {
	AppName       string
	ID            uint32
	AppIcon       string
	Summary       string
	Body          string
	Actions       []string
	Hints         map[string]dbus.Variant
	ExpireTimeout int32
}
