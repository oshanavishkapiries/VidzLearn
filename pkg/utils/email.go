package utils

import (
	"bytes"
	"html/template"
	"path/filepath"
)

// SendOtpEmail loads the OTP email HTML template and fills it with the OTP value.
func SendOtpEmail(email, otp string) (string, error) {
	tmplPath := filepath.Join("pkg", "smtp", "otp_email.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return "", err
	}

	// Data to inject into the template
	data := struct {
		Email string
		Otp   string
	}{
		Email: email,
		Otp:   otp,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
