package services

import (
    "bytes"
    "encoding/json"
    "net/http"
    "os"
)

type EmailRequest struct {
    Sender struct {
        Email string `json:"email"`
    } `json:"sender"`
    To []struct {
        Email string `json:"email"`
    } `json:"to"`
    Subject string `json:"subject"`
    HtmlContent string `json:"htmlContent"`
}

func SendEmailHTML(to, subject, body string) error {

    apiKey := os.Getenv("BREVO_API_KEY")

    email := EmailRequest{}
    email.Sender.Email = os.Getenv("EMAIL_SENDER")
    email.Subject = subject
    email.HtmlContent = body

    email.To = append(email.To, struct {
        Email string `json:"email"`
    }{Email: to})

    jsonData, _ := json.Marshal(email)

    req, _ := http.NewRequest("POST", "https://api.brevo.com/v3/smtp/email", bytes.NewBuffer(jsonData))

    req.Header.Set("accept", "application/json")
    req.Header.Set("api-key", apiKey)
    req.Header.Set("content-type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    return nil
}