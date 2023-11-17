package service

import (
	"fmt"
	"os"

	"github.com/tedxub2023/internal/mainevent"
	m "github.com/tedxub2023/internal/ticket/service"
)

func sendTransactionDeclinedMail(tx mainevent.MainEvent) error {
	mail := m.NewMailClient()
	mail.SetSender("tedxuniversitasbrawijaya@gmail.com")
	mail.SetReciever(tx.Email)
	mail.SetSubject("Transaksi Ditolak")

	if err := mail.SetBodyHTMLDeclinedTransaction(); err != nil {
		return err
	}

	if err := mail.SendMail(); err != nil {
		return err
	}
	return nil
}

func sendSuccessTransactionMail(tx mainevent.MainEvent) error {
	mail := m.NewMailClient()
	mail.SetSender("tedxuniversitasbrawijaya@gmail.com")
	mail.SetReciever(tx.Email)
	mail.SetSubject("Tiket Memantik Baskara")

	mail.SetAttachFile(fmt.Sprintf("global/storage/ted/%s-%s.pdf", tx.Nama, tx.Type.String()))

	if err := mail.SetBodyHTMLSuccessTransaction(tx); err != nil {
		return err
	}

	if err := mail.SendMail(); err != nil {
		return err
	}
	return nil
}

func sendMailInformAdmin(tx mainevent.MainEvent) error {
	mail := m.NewMailClient()
	mail.SetSender("tedxuniversitasbrawijaya@gmail.com")
	email := os.Getenv("EMAIL_CEM")
	mail.SetReciever(email)
	mail.SetSubject("Pemberitahuan Pembelian Tiket")

	if err := mail.SetBodyHTMLInformAdmin(tx); err != nil {
		return err
	}

	if err := mail.SendMail(); err != nil {
		return err
	}
	return nil
}

func sendMainEventPendingMail(tx mainevent.MainEvent) error {
	mail := m.NewMailClient()
	mail.SetSender("tedxuniversitasbrawijaya@gmail.com")
	mail.SetReciever(tx.Email)
	mail.SetSubject("Konfirmasi Pembelian Tiket")

	if err := mail.SetBodyHTMLMainEventPendingMail(); err != nil {
		return err
	}

	if err := mail.SendMail(); err != nil {
		return err
	}
	return nil
}
