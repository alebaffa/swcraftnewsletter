package private

func GetParams() (string, string, string, string) {
	email := "alessandro.baffa@gmail.com"
	password := "fake"
	smtpServer := "smtp.gmail.com"
	port := "587"
	return email, password, smtpServer, port
}
