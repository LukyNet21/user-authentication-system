package utils

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var D *gomail.Dialer

func InitEmail() {
	server := os.Getenv("EMAIL_SERVER_ADDRESS")
	port, err := strconv.Atoi(os.Getenv("EMAIL_SERVER_PORT"))
	if err != nil {
		panic("Invalid port")
	}
	username := os.Getenv("EMAIL_SERVER_USERNAME")
	password := os.Getenv("EMAIL_SERVER_PASSWORD")
	D = gomail.NewDialer(server, port, username, password)

}

func SendTestEmail(to string) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_SERVER_USERNAME"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Test email!")
	m.SetBody("text/html", "This is a test email!")

	if err := D.DialAndSend(m); err != nil {
		panic(err)
	}
}
