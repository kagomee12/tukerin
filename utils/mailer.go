package utils

import (
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

func SendEmail(to string, subject string, body string) error {

	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASS")
	host := os.Getenv("EMAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))


	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, from, password)
	
	if err:= d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}