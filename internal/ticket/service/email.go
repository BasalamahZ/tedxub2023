package service

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"text/template"

	"github.com/tedxub2023/internal/ticket"
	"gopkg.in/gomail.v2"
)

type Gomail struct {
	message *gomail.Message
	dialer  *gomail.Dialer
}

func NewMailClient() *Gomail {
	port, err := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	if err != nil {
		log.Fatalf("[tedxub2023-api-http] failed to convert smtp port: %s\n", err.Error())
	}

	return &Gomail{
		gomail.NewMessage(),
		gomail.NewDialer(
			os.Getenv("CONFIG_SMTP_HOST"),
			port,
			os.Getenv("CONFIG_AUTH_EMAIL"),
			os.Getenv("CONFIG_AUTH_PASSWORD"),
		)}
}

func (g *Gomail) SetSender(sender string) {
	g.message.SetHeader("From", sender)
}

func (g *Gomail) SetReciever(to ...string) {
	g.message.SetHeader("To", to...)
}

func (g *Gomail) SetSubject(subject string) {
	g.message.SetHeader("Subject", subject)
}

func (g *Gomail) SetBodyHTML(nameReciever string) error {
	var body bytes.Buffer
	path := "global/template/template.html"
	t, err := template.ParseFiles(path)
	if err != nil {
		return ticket.ErrParseBodyHTML
	}

	t.Execute(&body, struct {
		Nama string
	}{
		Nama: nameReciever,
	})
	g.message.SetBody("text/html", body.String())
	return nil
}

func (g *Gomail) SendMail() error {
	if err := g.dialer.DialAndSend(g.message); err != nil {
		return ticket.ErrSendEmail
	}
	return nil
}

func (g *Gomail) SetBodyHTMLMainEvent(totalTickets int, totalPrice string, dateTime string) error {
	var body bytes.Buffer
	path := "global/template/mainEvent.html"

	t, err := template.ParseFiles(path)
	if err != nil {
		return ticket.ErrParseBodyHTML
	}

	t.Execute(&body, struct {
		TotalTickets int
		TotalPrice   string
		DateTime     string
	}{
		TotalTickets: totalTickets,
		TotalPrice:   totalPrice,
		DateTime:     dateTime,
	})

	g.message.SetBody("text/html", body.String())
	return nil
}

func (g *Gomail) SetAttachFile(path string) {
	g.message.Attach(path)
}

func (g *Gomail) SetBodyHTMLPendingMail(name string, totalTickets int, totalPrice string, dateTime string) error {
	var body bytes.Buffer
	path := "global/template/pendingMail.html"

	t, err := template.ParseFiles(path)
	if err != nil {
		return ticket.ErrParseBodyHTML
	}

	t.Execute(&body, struct {
		Name         string
		TotalTickets int
		TotalPrice   string
		DateTime     string
	}{
		Name:         name,
		TotalTickets: totalTickets,
		TotalPrice:   totalPrice,
		DateTime:     dateTime,
	})

	g.message.SetBody("text/html", body.String())
	return nil
}
