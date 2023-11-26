package helper

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"

	"github.com/tedxub2023/internal/mainevent"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

func downloadImage(url string) (image.Image, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func PDF(tx mainevent.MainEvent) error {
	err := license.SetMeteredKey(os.Getenv("UNIDOC_LICENSE_API_KEY"))
	if err != nil {
		panic(err)
	}

	c := creator.New()

	for _, nomorTicket := range tx.NomorTiket {

		// c.SetPageMargins(30, 50, 100, 70)

		helvetica, _ := model.NewStandard14Font("Helvetica")

		img, err := c.NewImageFromFile("global/storage/ted/background.png")
		if err != nil {
			return err
		}

		img.ScaleToWidth(612.0)

		height := 612.0 * img.Height() / img.Width()
		c.SetPageSize(creator.PageSize{612, height})
		c.NewPage()
		img.SetPos(0, 0)
		c.Draw(img)

		// getting wrapper
		p := c.NewParagraph("Tiket Event")
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetMargins(-30, 30, 50, 0)
		p.SetColor(creator.ColorBlack)
		c.Draw(p)

		p = c.NewParagraph("Memantik Baskara | TEDxUniversitasBrawijaya")
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetMargins(-30, 30, 0, 0)
		p.SetLineHeight(1.5)
		p.SetColor(creator.ColorBlack)
		c.Draw(p)

		p = c.NewParagraph("3 Desember 2023")
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetMargins(-30, 30, 0, 0)
		p.SetLineHeight(1.5)
		p.SetColor(creator.ColorBlack)
		c.Draw(p)

		// baskara
		p = c.NewParagraph(nomorTicket)
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetMargins(-30, 30, 20, 0)
		p.SetColor(creator.ColorBlack)
		c.Draw(p)

		p = c.NewParagraph("Detail Audiens")
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetMargins(-30, 30, 0, 0)
		p.SetLineHeight(1.5)
		p.SetColor(creator.ColorBlack)
		c.Draw(p)

		// biodata
		name := fmt.Sprintf("Nama: %s", tx.Nama)
		p = c.NewParagraph(name)
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetMargins(-30, 30, 20, 0)
		p.SetColor(creator.ColorBlack)
		c.Draw(p)

		identityNumber := fmt.Sprintf("Nomor Identitas: %s", tx.NomorIdentitas)
		p = c.NewParagraph(identityNumber)
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetMargins(-30, 30, 0, 0)
		p.SetColor(creator.ColorBlack)
		c.Draw(p)

		email := fmt.Sprintf("Email: %s", tx.Email)
		p = c.NewParagraph(email)
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetMargins(-30, 30, 0, 0)
		p.SetColor(creator.ColorBlack)
		c.Draw(p)

		phone := fmt.Sprintf("No Telpon: %s", tx.NomorTelepon)
		p = c.NewParagraph(phone)
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetMargins(-30, 30, 0, 0)
		p.SetColor(creator.ColorBlack)
		c.Draw(p)

		ticketType := fmt.Sprintf("Jenis Tiket: %s", "Normal Sale")
		p = c.NewParagraph(ticketType)
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetMargins(-30, 30, 0, 0)
		p.SetColor(creator.ColorBlack)
		c.Draw(p)

		// barcode
		url := fmt.Sprintf(os.Getenv("URL_QRCODE"), tx.ID, nomorTicket)
		image, err := downloadImage("https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=" + url)
		if err != nil {
			return err
		}
		barcode, err := c.NewImageFromGoImage(image)
		if err != nil {
			return err
		}

		barcode.ScaleToWidth(170)

		barcode.SetPos(410, 220)
		c.Draw(barcode)

		// tatacara
		p = c.NewParagraph("Tata Cara Penukaran Tiket")
		p.SetFont(helvetica)
		p.SetMargins(-30, -30, 30, 0)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetColor(creator.ColorBlack)
		p.SetTextAlignment(creator.TextAlignmentJustify)
		c.Draw(p)

		p = c.NewParagraph("1. Silahkan kunjungi entrance gate dan tunjukan unique barcode yang telah kamu dapatkanuntuk di-scan oleh panitia yang bertugas;")
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetMargins(-30, -30, 0, 0)
		p.SetLineHeight(1.5)
		p.SetColor(creator.ColorBlack)
		p.SetTextAlignment(creator.TextAlignmentJustify)
		c.Draw(p)

		p = c.NewParagraph("2. Setelah unique barcode terverifikasi, kamu akan mendapatkan wristband dan juga TEDx Kit;")
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetColor(creator.ColorBlack)
		p.SetMargins(-30, -30, 0, 0)
		p.SetTextAlignment(creator.TextAlignmentJustify)
		c.Draw(p)

		p = c.NewParagraph("3. Panitia yang bertugas akan mengarahkanmu untuk menempati kursi yang telah tersedia;")
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetColor(creator.ColorBlack)
		p.SetMargins(-30, -30, 0, 0)
		p.SetTextAlignment(creator.TextAlignmentJustify)
		c.Draw(p)

		p = c.NewParagraph("4. Jika tiket kamu digunakan oleh orang lain, orang tersebut harus menunjukan bukti berupa foto kartu identitas (nama yang tertera pada kartu identitas harus sesuai dengan yang tertera pada e-ticket).")
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetColor(creator.ColorBlack)
		p.SetMargins(-30, -30, 0, 0)
		p.SetTextAlignment(creator.TextAlignmentJustify)
		c.Draw(p)

		p = c.NewParagraph("5. Penukaran tiket hanya dapat dilakukan pada sesi open gate yakni pukul 09.00 - 10.00 WIB. Jika audiens datang setelah sesi tersebut berakhir, maka otomatis tiket yang dimiliki audiens akan hangus dan audiens dilarang untuk memasuki venue acara.")
		p.SetFont(helvetica)
		p.SetFontSize(14)
		p.SetLineHeight(1.5)
		p.SetColor(creator.ColorBlack)
		p.SetMargins(-30, -30, 0, 0)
		p.SetTextAlignment(creator.TextAlignmentJustify)
		c.Draw(p)
	}

	path := fmt.Sprintf("global/storage/ted/%s-%s.pdf", tx.Nama, tx.Type.String())
	err = c.WriteToFile(path)
	if err != nil {
		log.Println("Write file error:", err)
	}

	return nil
}
