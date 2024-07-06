package gomail

import (
	"errors"
	"fmt"
	"gomodule/pkg/vip"
	"gomodule/pkg/zlog"
	"net/smtp"
)

type MailConf struct {
	Port     string
	Host     string
	Sender   string
	Password string
}

func New() *MailConf {
	vipp, err := vip.New().App()
	if err != nil {
		zlog.Error(errors.New(fmt.Sprintf("gomail service: %v", err)))
	}
	return &MailConf{
		Port:     vipp.EmailPort,
		Host:     vipp.EmailHost,
		Sender:   vipp.EmailSender,
		Password: vipp.EmailPassword,
	}
}

func (m *MailConf) Send(dest, subject, msg string) {
	message := []byte(fmt.Sprintf("To: %s\r\n", dest) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"\r\n" +
		msg)
	auth := smtp.PlainAuth("", m.Sender, m.Password, m.Host)
	err := smtp.SendMail(m.Host+":"+m.Port, auth, m.Sender, []string{dest}, message)
	if err != nil {
		zlog.Error(err)
	}
	zlog.Info(nil, "Email has been sent")
}
