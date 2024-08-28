package infrastructres

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/hunderaweke/spher/config"
)

//go:embed templates/email_template.html
var emailTemplate embed.FS

func SendEmail(subject, to string, data interface{}) error {
	tmpl, err := template.ParseFS(emailTemplate, "templates/email_template.html")
	if err != nil {
		return err
	}
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("error rendering template", err)
	}
	config, err := config.LoadConfig()
	if err != nil {
		return err
	}
	from := "hundera.awoke@a2sv.org"
	key := config.Email.Key
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	message := []byte("Subject: " + subject + "\n" + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + body.String())
	auth := smtp.PlainAuth("", from, key, host)
	if err = smtp.SendMail(address, auth, from, []string{to}, message); err != nil {
		return err
	}
	return nil
}
