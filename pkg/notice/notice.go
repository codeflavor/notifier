package notice

import (
	"github.com/esiqveland/notify"
	"github.com/godbus/dbus"
)

const (
	// defaultTimeOut is the default time out of the pop up.
	defaultTimeOut = 5000
)

var (
	defaultActions = []string{"Dismiss", "cancel"}
)

// Notifier holds information regarding a dbus session used to send
// notifications.
type Notifier struct {
	conn *dbus.Conn
}

func newNotifier(appName, icon, iconName, summary, body string) *notify.Notification {
	return &notify.Notification{
		AppName: appName,
		// NOTE: check this out
		ReplacesID: uint32(0),
		AppIcon:    iconName,
		Summary:    summary,
		Body:       body,
		Actions:    defaultActions,
	}
}

func (n *Notifier) newSession() error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	n.conn = conn
	return nil
}

// NotifyUser creates a new notification and sends it to the user.
func (n *Notifier) NotifyUser(appName, icon, iconName, summary, body string, timeOut int) error {
	return nil
}
