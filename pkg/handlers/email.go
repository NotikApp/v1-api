package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func sendEmail(email, url, name string) error {

	from := "notik.fun@gmail.com"
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

	t, err := template.ParseFiles("pkg\\handlers\\index.html")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s, verify your account \n%s\n\n", name, mimeHeaders)))

	t.Execute(&body, struct {
		Url  string
		Name string
	}{
		Name: name,
		Url:  url,
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent!")
	return nil
}
