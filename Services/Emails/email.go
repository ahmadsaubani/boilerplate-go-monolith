package Emails

import (
	"boilerplate-go/Libraries/Helpers"
	"context"
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func (e EmailSetting) dialMail() *gomail.Dialer {
	dialer := gomail.NewDialer(
		e.Config.Host,
		e.Config.Port,
		e.Config.User,
		e.Config.Password,
	)

	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	return dialer
}

func (e EmailSetting) Send(ctx context.Context, config *gomail.Message) error {
	if err := e.dialMail().DialAndSend(config); err != nil {
		Helpers.LogErrorCritical(ctx, err)
		return err
	}
	return nil
}

func (e EmailSetting) Set(to []string, subject, body string) *gomail.Message {
	mailconfig := gomail.NewMessage()
	mailconfig.SetHeader("From", e.Config.EmailFrom)
	mailconfig.SetHeader("To", to...)
	mailconfig.SetHeader("Subject", subject)
	mailconfig.SetBody("text/html", body)
	return mailconfig
}
