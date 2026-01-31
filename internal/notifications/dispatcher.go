package notifications

// Dispatches dufferent notifcations types
import (
	"fmt"

	"github.com/KostyaBagr/duple-duple/internal/backup"
	cfg "github.com/KostyaBagr/duple-duple/internal/config"
	"github.com/KostyaBagr/duple-duple/internal/notifications/mail"
	"github.com/KostyaBagr/duple-duple/internal/utils"
)


func NotificationDumpDispatcher(receiver string, dumpStats backup.DumpFileStats) {
	isEmptyNotificationCfg := utils.IsEmpty(cfg.AppConfig.Notifications)

	if !isEmptyNotificationCfg {
		isEmptyEmailCfg := utils.IsEmpty(cfg.AppConfig.Notifications.Email)

		if !isEmptyEmailCfg {
			subject := fmt.Sprintf("Your %s dump was created", dumpStats.Dbms)
			mail.Sender(receiver, subject, dumpStats.String())
		}

	} else {
		fmt.Println("Email config is null. If you want to send emails, please check config.toml")
	}
}
