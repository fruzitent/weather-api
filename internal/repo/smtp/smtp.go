package smtp

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

type Config struct {
	From     string
	Host     string
	Password string
	Port     int
	Username string
}

func SendMail(config *Config, to string, subject string, body string) error {
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port), &tls.Config{
		ServerName: config.Host,
	})
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, config.Host)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	if err := client.Auth(auth); err != nil {
		return err
	}

	if err := client.Mail(config.From); err != nil {
		return err
	}

	if err := client.Rcpt(to); err != nil {
		return err
	}

	headers := make(map[string]string)
	headers["From"] = config.From
	headers["To"] = to
	headers["Subject"] = subject

	msg := ""
	for k, v := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + body

	wc, err := client.Data()
	if err != nil {
		return err
	}
	if _, err := wc.Write([]byte(msg)); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	if err := client.Quit(); err != nil {
		return err
	}

	return nil
}
