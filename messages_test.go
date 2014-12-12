package oshirase_test

import (
	"testing"

	"github.com/godbus/dbus"
	"github.com/pocke/oshirase"
)

var server *oshirase.Server
var client *dbus.Object

const (
	name    = "ServerName"
	vendor  = "ServerVendor"
	version = "ServerVersion"
)

func TestGetCapabilities(t *testing.T) {
	call := client.Call("org.freedesktop.Notifications.GetCapabilities", 0)
	if call.Err != nil {
		t.Fatal(call.Err)
	}

	if v, ok := call.Body[0].([]string); !ok {
		t.Errorf("GetCapabilities should return []string. But got %s", v)
	}
}

func TestGetServerInformation(t *testing.T) {
	call := client.Call("org.freedesktop.Notifications.GetServerInformation", 0)
	if call.Err != nil {
		t.Fatal(call.Err)
	}

	if call.Body[0] != name {
		t.Error("GetServerInformation should return Name")
	}
	if call.Body[1] != vendor {
		t.Error("GetServerInformation should return Vendor")
	}
	if call.Body[2] != version {
		t.Error("GetServerInformation should return version")
	}
	if call.Body[3] != oshirase.SpecVersion {
		t.Error("GetServerInformation should return SpecVersion")
	}
}

func init() {
	s, err := oshirase.NewServer(name, vendor, version)
	if err != nil {
		panic(err)
	}
	server = s

	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	client = obj
}
