package oshirase

import (
	"fmt"

	"github.com/godbus/dbus"
)

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
