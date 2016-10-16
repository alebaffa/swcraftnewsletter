package mail

import "github.com/alebaffa/swcraftnewsletter/utils"

func Send(link string) error {
	email := utils.ReadCredentials().Mail
	password := utils.ReadCredentials().Password

	mailSender := NewSender(
		Config{Username: email, Password: password, ServerHost: "smtp.gmail.com", ServerPort: "587", SenderAddr: ""})

	return mailSender.SendMail([]string{email}, []byte("To: "+email+"\r\n"+
		"Subject: New link for Software Craftsmanship Newsletter\r\n"+
		"\r\n"+
		link))
}
