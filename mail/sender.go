package mail

import "os"

func Send(link string) error {
	email := os.Getenv("G_EMAIL")
	password := os.Getenv("G_PASSW")

	mailSender := NewSender(
		Config{Username: email, Password: password, ServerHost: "smtp.gmail.com", ServerPort: "587", SenderAddr: ""})

	return mailSender.SendMail([]string{email}, []byte("To: "+email+"\r\n"+
		"Subject: New link for Software Craftsmanship Newsletter\r\n"+
		"\r\n"+
		link))
}
