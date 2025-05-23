package smtp

import (
	"bytes"
	"embed"
	"html/template"
	"net/mail"
	"net/url"
	"time"

	"git.fruzit.pp.ua/weather/api/internal/lib/smtp"
	"git.fruzit.pp.ua/weather/api/pkg/user/domain/entity"
	"git.fruzit.pp.ua/weather/api/pkg/user/port"
	entityWeather "git.fruzit.pp.ua/weather/api/pkg/weather/domain/entity"
)

//go:embed template/*.html.tmpl
var Template embed.FS

type Smtp struct {
	Config *smtp.Config
}

func NewSmtp(config *smtp.Config) *Smtp {
	return &Smtp{config}
}

var _ port.Notification = (*Smtp)(nil)

func (a *Smtp) SendConfirmation(user entity.User) error {
	addr, err := mail.ParseAddress(user.Mail.Address)
	if err != nil {
		return err
	}

	// TODO: spaghetti
	uri, err := url.Parse("https://weather.fruzit.pp.ua/api/confirm/1234")
	if err != nil {
		return err
	}

	body := &bytes.Buffer{}

	tmpl, err := template.ParseFS(Template, "template/confirm.html.tmpl")
	if err != nil {
		return err
	}
	if err := tmpl.ExecuteTemplate(body, "confirm.html.tmpl", map[string]string{
		"Confirm": uri.String(),
	}); err != nil {
		return err
	}

	return smtp.SendMail(a.Config, addr, "Confirm Subscription", body.Bytes())
}

func (a *Smtp) SendWeatherReport(user entity.User, report entityWeather.Report) error {
	addr, err := mail.ParseAddress(user.Mail.Address)
	if err != nil {
		return err
	}

	// TODO: spaghetti
	uri, err := url.Parse("https://weather.fruzit.pp.ua/api/unsubscribe/1234")
	if err != nil {
		return err
	}

	body := &bytes.Buffer{}

	tmpl, err := template.ParseFS(Template, "template/report.html.tmpl")
	if err != nil {
		return err
	}
	if err := tmpl.ExecuteTemplate(body, "report.html.tmpl", map[string]string{
		"Date":        time.Unix(report.CreatedAt, 0).Format(time.RFC3339),
		"Description": report.Forecast.Description,
		"Humidity":    report.Forecast.Humidity.String(),
		"Temperature": report.Forecast.Temperature.String(),
		"Unsubscribe": uri.String(),
	}); err != nil {
		return err
	}

	return smtp.SendMail(a.Config, addr, "Weather Report", body.Bytes())
}
