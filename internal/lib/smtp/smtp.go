package smtp

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"mime"
	"net/mail"
	"net/smtp"
	"strings"
	"time"
)

const CRLF string = "\r\n"

type Config struct {
	From     mail.Address
	Host     string
	Password string
	Port     int
	Username string
}

func SendMail(config *Config, to *mail.Address, subject string, body []byte) error {
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

	if err := client.Mail(config.From.Address); err != nil {
		return err
	}

	if err := client.Rcpt(to.Address); err != nil {
		return err
	}

	noreply, err := mail.ParseAddress(fmt.Sprintf("no-reply@%s", config.Host))
	if err != nil {
		return err
	}

	headers := make(map[string]string)
	headers["Content-Transfer-Encoding"] = "base64"
	headers["Content-Type"] = "text/html; charset=utf-8"
	headers["Date"] = time.Now().Format(time.RFC1123Z)
	headers["From"] = config.From.String()
	// TODO: POST request
	// headers["List-Unsubscribe"] = ""
	// headers["List-Unsubscribe-Post"] = "List-Unsubscribe=One-Click"
	headers["MIME-Version"] = "1.0"
	headers["Reply-To"] = noreply.String()
	headers["Subject"] = mime.QEncoding.Encode("utf-8", subject)
	headers["To"] = to.String()

	msg := ""
	for k, v := range headers {
		header := fmt.Sprintf("%s: %s", k, v)
		if strings.ContainsAny(header, CRLF) {
			return fmt.Errorf("RFC 5321 violation %s", header)
		}
		msg += header + CRLF
	}
	msg += CRLF + base64.StdEncoding.EncodeToString(body) + CRLF

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
