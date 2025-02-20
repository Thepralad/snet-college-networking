package services
import(
    "net/smtp"
    "os"
    "github.com/joho/godotenv"
)

func SendEmail(to []string, subject string, body string) error{
    godotenv.Load()
    auth := smtp.PlainAuth(
        "",
        os.Getenv("FROM_EMAIL"),
        os.Getenv("FROM_EMAIL_PASSWORD"),
        os.Getenv("FROM_EMAIL_SMTP"),
    )
    message := "Subject: " + subject + "\r\n" +
               "Content-Type: text/html; charset=\"utf-8\"\r\n" +
               "\r\n" + 
               body
    return smtp.SendMail(os.Getenv("SMTP_ADDR"), auth, os.Getenv("FROM_EMAIL"), to , []byte(message))
}
