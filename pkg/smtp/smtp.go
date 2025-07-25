package smtp

import (
	"fmt"
	"os"

	"github.com/Cenzios/pf-backend/pkg/logger"
	mail "github.com/go-mail/mail/v2"
)

var (
	smtpHost      string
	smtpPort      int
	smtpUsername  string
	smtpPassword  string
	smtpFromEmail string
	smtpFromName  string
)

// Init loads SMTP config from environment variables
func Init() error {
	smtpHost = os.Getenv("SMTP_HOST")
	smtpPort = 587 // default
	if port := os.Getenv("SMTP_PORT"); port != "" {
		fmt.Sscanf(port, "%d", &smtpPort)
	}
	smtpUsername = os.Getenv("SMTP_USERNAME")
	smtpPassword = os.Getenv("SMTP_PASSWORD")
	smtpFromEmail = os.Getenv("SMTP_FROM_EMAIL")
	smtpFromName = os.Getenv("SMTP_FROM_NAME")

	if smtpHost == "" || smtpUsername == "" || smtpPassword == "" || smtpFromEmail == "" || smtpFromName == "" {
		return fmt.Errorf("SMTP config missing required fields")
	}
	logger.Info.Println("✅ Connected to SMTP")
	return nil
}

// SendMail sends an email using the loaded SMTP config
func SendMail(to, subject, body string) error {
	m := mail.NewMessage()
	m.SetAddressHeader("From", smtpFromEmail, smtpFromName)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := mail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	return d.DialAndSend(m)
}
