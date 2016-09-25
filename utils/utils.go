package utils

import (
	"net/smtp"

	"github.com/alebaffa/newsletter-web/private"
)

func SendEmail(link string) error {
	email, password, smtpServer, port := private.GetParams()

	auth := smtp.PlainAuth(
		"",
		email,
		password,
		smtpServer,
	)
	return smtp.SendMail(
		smtpServer+":"+port,
		auth,
		"sender@email.com",
		[]string{email},
		[]byte("To: "+email+"\r\n"+
			"Subject: New link for Software Craftsmanship Newsletter\r\n"+
			"\r\n"+
			link),
	)
}
