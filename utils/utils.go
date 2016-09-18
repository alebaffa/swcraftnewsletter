package utils

import (
	"html/template"
	"net/http"
	"net/smtp"

	"github.com/alebaffa/newsletter-web/private"
)

func ExecuteTemplate(templateType string, res http.ResponseWriter) {
	templateType = templateType + ".html"
	templates := template.Must(template.ParseGlob("assets/*.html"))
	if err := templates.ExecuteTemplate(res, templateType, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

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
