package notifications

// Dispatches dufferent notifcations types
import (
	cfg "github.com/KostyaBagr/duple-duple/internal/config"
	"github.com/KostyaBagr/duple-duple/internal/notifications/mail"
	"github.com/KostyaBagr/duple-duple/internal/utils"
)

// TODO: rewriter pasta code
func NotificationDispather(receiver, subject, body string) {
	isNotificationCfg := utils.IsEmpty(cfg.AppConfig.Notifications)

	if isNotificationCfg {
		isEmailCfg := utils.IsEmpty(cfg.AppConfig.Notifications.Email)

		if isEmailCfg {
			// TODO: add more statistics to email body (taken time, size of files, etc)
			mail.Sender(receiver, subject, body)
		}

	}
}
