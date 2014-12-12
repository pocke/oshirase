package oshirase

import "github.com/godbus/dbus"

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
	SpecVersion string = "1.2"
)

// action-icons, actions, body, body-hyperlinks, body-images, body-markup, icon-multi, icon-static, persistence, sound
func (m messages) GetCapabilities() ([]string, *dbus.Error) {
	// TODO: select return variables
	return []string{"body"}, nil
}

func (m messages) Notify(appName string, replacesID uint32, appIcon, summary, body string, actions []string, hints map[string]dbus.Variant, expireTimeout int32) (id uint32, err *dbus.Error) {
	if replacesID == 0 {
		id = <-m.server.notifyID
	} else {
		id = replacesID
	}

	a := &NotifyArg{
		AppName:       appName,
		ID:            id,
		AppIcon:       appIcon,
		Summary:       summary,
		Body:          body,
		Actions:       actions,
		Hints:         hints,
		ExpireTimeout: expireTimeout,
	}

	if m.server.onNotify != nil {
		m.server.onNotify(a)
	}

	// XXX: error is nil, That correct?
	return id, nil
}

func (m messages) CloseNotification(id uint32) *dbus.Error {
	if m.server.onCloseNotification == nil {
		return nil
	}

	ok := m.server.onCloseNotification(id)
	// when id dose not exists
	if !ok {
		// XXX: That correct?
		return dbus.NewError("org.freedesktop.Notifications.CloseNotification", nil)
	}

	m.server.conn.Emit("/org/freedesktop/Notifications/NotificationClosed", "org.freedesktop.Notifications.NotificationClosed", id, CloseReasonCallToCloseNotification)
	return nil
}

func (m messages) GetServerInformation() (name, vendor, version, specVer string, err *dbus.Error) {
	return m.server.name, m.server.vendor, m.server.version, SpecVersion, nil
}
