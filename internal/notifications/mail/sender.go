// Package for sending emails
package mail

import (
	"errors"
	"fmt"

	cfg "github.com/KostyaBagr/duple-duple/internal/config"
	gomail "gopkg.in/mail.v2"
)

//  TODO: add retries if smtp server is down
// Sends regular email in text format (TODO: add html rendering)
func Sender(receiver, subject, body string) error {
	message := gomail.NewMessage()

	message.SetHeader("From", cfg.AppConfig.Notifications.Email.Sender)
	message.SetHeader("To", receiver)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	dialer := gomail.NewDialer(
		cfg.AppConfig.Notifications.Email.SmtpServer,
		cfg.AppConfig.Notifications.Email.Port,
		cfg.AppConfig.Notifications.Email.Sender,
		cfg.AppConfig.Notifications.Email.Password,
	)

	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error:", err)
		return errors.New("error during sending email")
	}
	fmt.Printf("Email sent successfully to %s!\n", receiver)
	return nil
}
