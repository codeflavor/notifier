package notice

import (
	"github.com/esiqveland/notify"
)

func newNotifier(icon, iconName, summary, body string, timeOut int32) *notify.Notification {
	return &notify.Notification{}
}

// NotifyUser creates a new notification and sends it to the user.
func NotifyUser(msg, iconName, summary, body string, timeOut int32) error {
	return nil
}
