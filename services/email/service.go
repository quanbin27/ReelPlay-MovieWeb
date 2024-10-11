package email

import (
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	smtpHost string
	smtpPort int
	username string
	password string
	from     string
}

func NewEmailService(host string, port int, username, password, from string) *EmailService {
	return &EmailService{
		smtpHost: host,
		smtpPort: port,
		username: username,
		password: password,
		from:     from,
	}
}

func (s *EmailService) SendResetPasswordEmail(to, resetLink string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.from)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Password Reset Request")

	// Body email
	body := "Please click the link below to reset your password:\n" + resetLink
	mailer.SetBody("text/plain", body)

	// Thiết lập cấu hình SMTP
	dialer := gomail.NewDialer(s.smtpHost, s.smtpPort, s.username, s.password)

	// Gửi email
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}
