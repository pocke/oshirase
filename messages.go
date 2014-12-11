package oshirase

import (
	"fmt"

	"github.com/godbus/dbus"
)

type messages struct {
	server *Server
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

// action-icons, actions, body, body-hyperlinks, body-images, body-markup, icon-multi, icon-static, persistence, sound
func (m messages) GetCapabilities() ([]string, *dbus.Error) {
	// TODO: select return variables
	return []string{"body"}, nil
}

func (m messages) Notify(appName string, replacesID uint32, appIcon, summary, body string, actions []string, hints map[string]dbus.Variant, expireTimeout int32) (uint32, *dbus.Error) {
	// TODO: result
	// TODO: Do notify
	fmt.Println(summary)
	fmt.Println(body)
	return 1, nil
}

func (m messages) CloseNotification(id uint32) *dbus.Error {
	// TODO: Do delete Notification, and return delete success?
	//       if notification dosen't exists, return empty dbus-error
	m.server.conn.Emit("/org/freedesktop/Notifications/NotificationClosed", "org.freedesktop.Notifications.NotificationClosed", id, CloseReasonCallToCloseNotification)
	return nil
}

func (m messages) GetServerInformation() (name, vendor, version, specVer string, err *dbus.Error) {
	return m.server.name, m.server.vendor, m.server.version, specVersion, nil
}
