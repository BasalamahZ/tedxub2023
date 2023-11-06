package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/tedxub2023/global/helper"
)

type merchHandler struct {
}

func (h *merchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGetAllMerch(w, r)
	default:
	}
}

func (h *merchHandler) handleGetAllMerch(w http.ResponseWriter, r *http.Request) {
	// add timeout to context
	ctx, cancel := context.WithTimeout(r.Context(), 2000*time.Millisecond)
	defer cancel()

	var (
		err        error           // stores error in this handler
		source     string          // stores request source
		resBody    []byte          // stores response body to write
		statusCode = http.StatusOK // stores response status code
	)

	// write response
	defer func() {
		// error
		if err != nil {
			log.Printf("[Merch HTTP][handleGetAllMerch] Failed to get merch. Source: %s, Err: %s\n", source, err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan []MerchHTTP, 1)
	errChan := make(chan error, 1)

	go func() {
		data := []MerchHTTP{
			{
				ID:        1,
				Nama:      "Kaos (Ring Tee)",
				Harga:     "Rp129.900",
				Deskripsi: "Setiap individu merupakan pribadi yang unik. Keunikan ini dapat kamu wujudkan melalui bahasa rupa yang imajinatif!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712259/MERCH%20%2B%20BG/Artboard_1_wtscgn.png",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        2,
				Nama:      "Kaos (Oversize)",
				Harga:     "Rp139.900",
				Deskripsi: "Manifestasikan intuisi dirimu sebagai insan yang bergelora.",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712260/MERCH%20%2B%20BG/Artboard_2_sdp4gr.png",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        3,
				Nama:      "Totebag",
				Harga:     "Rp74.900",
				Deskripsi: "Adakalanya kreativitas membawa dirimu hanyut dalam dimensi yang tak terlupakan.",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712258/MERCH%20%2B%20BG/TOTE_BAG_p0eivf.png",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        4,
				Nama:      "Lanyard",
				Harga:     "Rp25.000",
				Deskripsi: "Personalisasikan dan ekspresikan dirimu di antara eksklusivitas yang mendominasi!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712256/MERCH%20%2B%20BG/LANYARD_okcqeo.png",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        5,
				Nama:      "Sticker Pack",
				Harga:     "Rp25.000",
				Deskripsi: "Eksplorasi imajinatif melahirkan keunikan dengan nilai tak terbatas.",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712260/MERCH%20%2B%20BG/stickerpack_z552r2.png",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        6,
				Nama:      "Korek",
				Harga:     "Rp15.000",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712260/MERCH%20%2B%20BG/LIGHTER_pkr0rf.png",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        7,
				Nama:      "Hoodie",
				Harga:     "Rp220.000",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699266073/MERCH%20Batch%202/Rectangle_379-4_lkrzz8.jpg",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        8,
				Nama:      "Oversized T-shirt",
				Harga:     "Rp120.000",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699266072/MERCH%20Batch%202/Rectangle_379_facwhb.jpg",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        9,
				Nama:      "T-shirt",
				Harga:     "Rp100.000",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699266071/MERCH%20Batch%202/Rectangle_379-1_cssfqr.jpg",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        10,
				Nama:      "Sticker Pack",
				Harga:     "Rp15.000",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699266073/MERCH%20Batch%202/Rectangle_379-2_cd1teh.jpg",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        11,
				Nama:      "Korek",
				Harga:     "Rp15.000",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699266074/MERCH%20Batch%202/Rectangle_379-3_fjhihf.jpg",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        12,
				Nama:      "T-Shirt (Ring Tee)",
				Harga:     "Rp129.900",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699265957/MERCH%20Batch%202/Ring_Tee_wqfvry.jpg",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        13,
				Nama:      "T-Shirt (Oversize)",
				Harga:     "Rp139.900",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699265956/MERCH%20Batch%202/Oversize_znydzz.jpg",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        14,
				Nama:      "Lanyard",
				Harga:     "Rp25.000",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699265956/MERCH%20Batch%202/Lanyard_czbfbg.jpg",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        15,
				Nama:      "Korek",
				Harga:     "Rp15.000",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699265959/MERCH%20Batch%202/Korek_cdr9fw.jpg",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        16,
				Nama:      "Totebag",
				Harga:     "Rp74.900",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699285412/MERCH%20Batch%202/Group_324_tyxo1y.jpg",
				Link:      "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        17,
				Nama:      "Bundling 1",
				Harga:     "Rp125.000",
				Deskripsi: "Penawaran Menarik untukmu! Merchandise istimewa ini dapat kamu miliki dengan harga yang spesial!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699285972/MERCH%20Batch%202/bundling1-thumbnail_u9cwjv.png",
				Detail: []string{
					"https://res.cloudinary.com/dpcwbnax4/image/upload/v1699266243/MERCH%20Batch%202/Bundling_1_fkgbkd.png",
				},
				Link: "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        18,
				Nama:      "Bundling 2",
				Harga:     "Rp145.000",
				Deskripsi: "Penawaran Menarik untukmu! Merchandise istimewa ini dapat kamu miliki dengan harga yang spesial!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699285972/MERCH%20Batch%202/bundling2-thumbnail_nzju9y.png",
				Detail: []string{
					"https://res.cloudinary.com/dpcwbnax4/image/upload/v1699266246/MERCH%20Batch%202/Bundling_2_coosh1.png",
				},
				Link: "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
			{
				ID:        19,
				Nama:      "Extra Bundling",
				Harga:     "Rp225.000",
				Deskripsi: "Penawaran Menarik untukmu! Merchandise istimewa ini dapat kamu miliki dengan harga yang spesial!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1699285972/MERCH%20Batch%202/bundling3-thumbnail_n89hcu.png",
				Detail: []string{
					"https://res.cloudinary.com/dpcwbnax4/image/upload/v1699266236/MERCH%20Batch%202/Extra_Bundling_thmlbl.png",
				},
				Link: "https://forms.gle/rXEmJSnip4w5kRgd6",
			},
		}

		resChan <- data
	}()

	// wait and handle main go routine
	select {
	case <-ctx.Done():
		statusCode = http.StatusGatewayTimeout
		err = errRequestTimeout
	case err = <-errChan:
	case data := <-resChan:
		// construct response data
		resBody, err = json.Marshal(helper.ResponseEnvelope{
			Data: data,
		})
	}
}
