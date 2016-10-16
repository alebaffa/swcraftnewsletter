package mail

import "net/smtp"

type Config struct {
	Username   string
	Password   string
	ServerHost string
	ServerPort string
	SenderAddr string
}

type Sender interface {
	SendMail(to []string, body []byte) error
}

func NewSender(conf Config) Sender {
	return &emailSender{conf, smtp.SendMail}
}

type emailSender struct {
	conf Config
	send func(string, smtp.Auth, string, []string, []byte) error
}

func (e *emailSender) SendMail(to []string, body []byte) error {
	addr := e.conf.ServerHost + ":" + e.conf.ServerPort
	auth := smtp.PlainAuth("", e.conf.Username, e.conf.Password, e.conf.ServerHost)
	return e.send(addr, auth, e.conf.SenderAddr, to, body)
}
