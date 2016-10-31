package mail

import (
	"net/smtp"
	"os"
)

//Config stores the information needed to send the email
type Config struct {
	Username   string
	Password   string
	ServerHost string
	ServerPort string
	SenderAddr string
}

//Sender is the interface to implement to send an email
type Sender interface {
	SendMail(to []string, body []byte) error
}

type emailSender struct {
	conf Config
	send func(string, smtp.Auth, string, []string, []byte) error
}

//NewSender returns a new instance of the emailSender struct declaring the default function used in production to send the email: smtp.SendMail. In the unit test this can be mocked.
func NewSender(conf Config) Sender {
	return &emailSender{conf, smtp.SendMail}
}

//emailSender implements Sender interface.
func (e *emailSender) SendMail(to []string, body []byte) error {
	addr := e.conf.ServerHost + ":" + e.conf.ServerPort
	auth := smtp.PlainAuth("", e.conf.Username, e.conf.Password, e.conf.ServerHost)
	return e.send(addr, auth, e.conf.SenderAddr, to, body)
}

//Send is the entry point for the web handler.
func Send(link string) error {
	email := os.Getenv("G_EMAIL")
	password := os.Getenv("G_PASSW")

	mailSender := NewSender(
		Config{
			Username:   email,
			Password:   password,
			ServerHost: "smtp.gmail.com",
			ServerPort: "587",
			SenderAddr: ""})

	return mailSender.SendMail(
		[]string{email},
		[]byte("To: "+email+"\r\n"+
			"Subject: New link for Software Craftsmanship Newsletter\r\n"+
			"\r\n"+
			link))
}
