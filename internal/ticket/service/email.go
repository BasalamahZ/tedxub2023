package service

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/leekchan/accounting"
	"github.com/tedxub2023/internal/mainevent"
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

func (g *Gomail) SetBodyHTMLMainEvent(totalTickets int, totalPrice string) error {
	var body bytes.Buffer
	path := "global/template/mainEvent.html"

	t, err := template.ParseFiles(path)
	if err != nil {
		return ticket.ErrParseBodyHTML
	}

	t.Execute(&body, struct {
		TotalTickets int
		TotalPrice   string
	}{
		TotalTickets: totalTickets,
		TotalPrice:   totalPrice,
	})

	g.message.SetBody("text/html", body.String())
	return nil
}

func (g *Gomail) SetAttachFile(path string) {
	g.message.Attach(path)
}

func (g *Gomail) SetBodyHTMLPendingMail() error {
	var body bytes.Buffer
	path := "global/template/pendingMail.html"

	htmlContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	body.Write(htmlContent)

	g.message.SetBody("text/html", body.String())
	return nil
}

func (g *Gomail) SetBodyHTMLDeclinedTransaction() error {
	var body bytes.Buffer

	path := "global/template/declinedTransaction.html"

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	body.Write(file)

	g.message.SetBody("text/html", body.String())
	return nil
}

func (g *Gomail) SetBodyHTMLSuccessTransaction(tx mainevent.MainEvent) error {
	path := "global/template/successTransaction.html"

	t, err := template.ParseFiles(path)
	if err != nil {
		return ticket.ErrParseBodyHTML
	}

	var body bytes.Buffer

	ac := accounting.Accounting{Symbol: "Rp", Precision: 0, Thousand: ".", Decimal: ","}
	totalPrice := ac.FormatMoney(tx.TotalHarga)

	t.Execute(&body, struct {
		TypeTickets  string
		Date         string
		TotalTickets int
		TotalPrice   string
	}{
		TypeTickets:  "Early Bird",
		Date:         "3 Desember 2023",
		TotalTickets: tx.JumlahTiket,
		TotalPrice:   totalPrice,
	})

	g.message.SetBody("text/html", body.String())
	return nil
}

func (g *Gomail) SetBodyHTMLMainEventPendingMail() error {
	var body bytes.Buffer

	path := "global/template/mainEventPendingMail.html"

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	body.Write(file)

	g.message.SetBody("text/html", body.String())
	return nil
}

func (g *Gomail) SetBodyHTMLInformAdmin(tx mainevent.MainEvent) error {
	path := "global/template/informEmail.html"

	t, err := template.ParseFiles(path)
	if err != nil {
		return ticket.ErrParseBodyHTML
	}

	var body bytes.Buffer

	t.Execute(&body, struct {
		Name  string
		Email string
		Date  string
	}{
		Name:  tx.Nama,
		Email: tx.Email,
		Date:  time.Now().Format("2006-01-02 15:04:05"),
	})

	g.message.SetBody("text/html", body.String())
	return nil
}
