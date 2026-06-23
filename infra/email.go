package infra

import (
	"net/smtp"

	"github.com/romanova09/bold-trunojoyo-scholarship-api/config"
)

type SMTPService struct {
	cfg *config.Config
}

func NewEmail(cfg *config.Config) *SMTPService {
	return &SMTPService{cfg: cfg}
}

func (s *SMTPService) SendEmail(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", s.cfg.SMTP.Username, s.cfg.SMTP.Password, s.cfg.SMTP.Host)

	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(s.cfg.SMTP.Host+":"+s.cfg.SMTP.Port, auth, s.cfg.SMTP.Username, to, msg)
	if err != nil {
		return err
	}
	return nil
}
