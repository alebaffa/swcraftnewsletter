package mail

import (
	"net/smtp"
	"testing"
)

func TestEmailSentSuccessful(t *testing.T) {
	//Given
	function, emailRecorder := mockSend()
	sender := &emailSender{send: function}
	to := "me@example.com"
	body := "Hello World"

	//When
	err := sender.SendMail([]string{to}, []byte(body))

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(emailRecorder.to[0]) != to {
		t.Errorf("wrong message addr.\n\nexpected: %s\n got: %s", to, emailRecorder.to)
	}
	if string(emailRecorder.body) != body {
		t.Errorf("wrong message body.\n\nexpected: %s\n got: %s", body, emailRecorder.body)
	}
}

//mockSend returns a mocked func to be used instead of smtp.SendMail, and a structure of the email to send
func mockSend() (func(string, smtp.Auth, string, []string, []byte) error, *emailRecorder) {
	r := new(emailRecorder)
	return func(addr string, a smtp.Auth, from string, to []string, body []byte) error {
		*r = emailRecorder{addr, a, from, to, body}
		return nil
	}, r
}

type emailRecorder struct {
	addr       string
	auth       smtp.Auth
	senderAddr string
	to         []string
	body       []byte
}
