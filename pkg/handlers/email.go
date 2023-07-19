package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

func sendEmail(email, url, name, urlDelete string) error {

	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("static" + string(os.PathSeparator) + "index.html")
	if err != nil {
		log.Println(err)
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s, verify your account \n%s\n\n", name, mimeHeaders)))

	t.Execute(&body, struct {
		Url       string
		Name      string
		UrlDelete string
	}{
		Name:      name,
		Url:       url,
		UrlDelete: urlDelete,
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
