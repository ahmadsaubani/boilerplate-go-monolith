package Emails

import (
	"boilerplate-go/Config"
	"context"

	"gopkg.in/gomail.v2"
)

type EmailSetting struct {
	Config *Config.EmailConfig
}

type EmailServices interface {
	dialMail() *gomail.Dialer
	Send(ctx context.Context, config *gomail.Message) (err error)
	Set(to []string, subject, body string) *gomail.Message
}
