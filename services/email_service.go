package services

import (
    "net/smtp"
    "os"
)

func SendEmailHTML(to, subject, body string) error {

    from := os.Getenv("EMAIL_SENDER")
    password := os.Getenv("EMAIL_PASSWORD")

    smtpHost := "smtp.gmail.com"
    smtpPort := "587"

    msg := "MIME-Version: 1.0\r\n" +
        "Content-type: text/html; charset=\"UTF-8\";\r\n" +
        "Subject: " + subject + "\r\n\r\n" +
        body

    auth := smtp.PlainAuth("", from, password, smtpHost)

    return smtp.SendMail(
        smtpHost+":"+smtpPort,
        auth,
        from,
        []string{to},
        []byte(msg),
    )
}