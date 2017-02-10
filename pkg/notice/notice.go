package notice

import (
	"github.com/esiqveland/notify"
)

type Notifier struct {
}

func newNotifier(icon, iconName, summary, body string, timeOut int32) *notify.Notification {
	return &notify.Notification{}
}
