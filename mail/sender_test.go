package mail

import (
	"net/smtp"
	"testing"
)

func TestEmail_SendSuccessful(t *testing.T) {
	function, emailRecorder := mockSend()

	sender := &emailSender{send: function}
	body := "Hello World"
	err := sender.SendMail([]string{"me@example.com"}, []byte(body))

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(emailRecorder.msg) != body {
		t.Errorf("wrong message body.\n\nexpected: %\n got: %s", body, emailRecorder.msg)
	}
}

func mockSend() (func(string, smtp.Auth, string, []string, []byte) error, *emailRecorder) {
	r := new(emailRecorder)
	return func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		*r = emailRecorder{msg}
		return nil
	}, r
}

type emailRecorder struct {
	msg []byte
}
