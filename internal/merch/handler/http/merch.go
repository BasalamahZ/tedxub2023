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
				Link:      "gform link",
			},
			{
				ID:        2,
				Nama:      "Kaos (Oversize)",
				Harga:     "Rp139.900",
				Deskripsi: "Manifestasikan intuisi dirimu sebagai insan yang bergelora.",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712260/MERCH%20%2B%20BG/Artboard_2_sdp4gr.png",
				Link:      "gform link",
			},
			{
				ID:        3,
				Nama:      "Totebag",
				Harga:     "Rp74.900",
				Deskripsi: "Adakalanya kreativitas membawa dirimu hanyut dalam dimensi yang tak terlupakan.",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712258/MERCH%20%2B%20BG/TOTE_BAG_p0eivf.png",
				Link:      "gform link",
			},
			{
				ID:        4,
				Nama:      "Lanyard",
				Harga:     "Rp25.000",
				Deskripsi: "Personalisasikan dan ekspresikan dirimu di antara eksklusivitas yang mendominasi!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712256/MERCH%20%2B%20BG/LANYARD_okcqeo.png",
				Link:      "gform link",
			},
			{
				ID:        5,
				Nama:      "Sticker Pack",
				Harga:     "Rp25.000",
				Deskripsi: "Eksplorasi imajinatif melahirkan keunikan dengan nilai tak terbatas.",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712260/MERCH%20%2B%20BG/stickerpack_z552r2.png",
				Link:      "gform link",
			},
			{
				ID:        6,
				Nama:      "Korek",
				Harga:     "Rp15.000",
				Deskripsi: "Lakukan ekspedisimu dengan semangat yang terus memantik!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712260/MERCH%20%2B%20BG/LIGHTER_pkr0rf.png",
				Link:      "gform link",
			},
			{
				ID:        7,
				Nama:      "Bundling 1",
				Harga:     "Rp199.900",
				Deskripsi: "Penawaran Menarik untukmu! Merchandise istimewa ini dapat kamu miliki dengan harga yang spesial!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712257/MERCH%20%2B%20BG/Artboard_4_smrl1k.png",
				Link:      "gform link",
			},
			{
				ID:        8,
				Nama:      "Bundling 2",
				Harga:     "Rp199.900",
				Deskripsi: "Penawaran Menarik untukmu! Merchandise istimewa ini dapat kamu miliki dengan harga yang spesial!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712257/MERCH%20%2B%20BG/Artboard_5_a7ozwc.png",
				Link:      "gform link",
			},
			{
				ID:        9,
				Nama:      "Extra Bundling",
				Harga:     "Rp54.900",
				Deskripsi: "Penawaran Menarik untukmu! Merchandise istimewa ini dapat kamu miliki dengan harga yang spesial!",
				Thumbnail: "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694712256/MERCH%20%2B%20BG/Artboard_6_kpiozl.png",
				Link:      "gform link",
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
