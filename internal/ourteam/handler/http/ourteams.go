package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/tedxub2023/global/helper"
)

type ourteamsHandler struct {
}

func (h *ourteamsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGetAllOurTeams(w, r)
	default:
	}
}

func (h *ourteamsHandler) handleGetAllOurTeams(w http.ResponseWriter, r *http.Request) {
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
			log.Printf("[OurTeam HTTP][handleGetAllOurTeams] Failed to get our team. Source: %s, Err: %s\n", source, err.Error())
			helper.WriteErrorResponse(w, statusCode, []string{err.Error()})
			return
		}
		// success
		helper.WriteResponse(w, resBody, statusCode, helper.JSONContentTypeDecorator)
	}()

	// prepare channels for main go routine
	resChan := make(chan []OurTeam, 1)
	errChan := make(chan error, 1)

	go func() {
		data := []OurTeam{
			{
				ID:     1,
				Divisi: "Organizer & Co-Organizer",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Nabyl Fadhlurrahman",
						Fakultas:  "Fakultas Ilmu Sosial dan Ilmu Politik",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694845919/TEDXUB2023%20Our%20Team/Tambahan/Nabyl_tibfxq.png",
						Instagram: "nabylf",
						LinkedIn:  "Nabyl Fadhlurrahman",
					},
					{
						Nama:      "Muhammad Nur Mi'raj AL-Qadrie",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694845948/TEDXUB2023%20Our%20Team/Tambahan/Minem_jlfgqy.png",
						Instagram: "qadrieeee",
						LinkedIn:  "Muhammad Nur Mi'raj AL-Qadrie",
					},
				},
				Volunteer: []string{},
			},
			{
				ID:     2,
				Divisi: "Budget Manager",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Reihansyah Ardhiya Amanullah",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609800/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00430_uufyrv.jpg",
						Instagram: "reihardhiya",
						LinkedIn:  "Reihansyah Ardhiya Amanullah",
					},
					{
						Nama:      "Kezia Agustina Hatiuran",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609801/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00419_caqxui.jpg",
						Instagram: "keziagustina",
						LinkedIn:  "keziagustina",
					},
				},
				Volunteer: []string{},
			},
			{
				ID:     3,
				Divisi: "Curator",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Pricilla Philia Br. Purba",
						Fakultas:  "Fakultas Ilmu Sosial dan Ilmu Politik",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609802/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00394_twbsfa.jpg",
						Instagram: "pricillaphiliaa",
						LinkedIn:  "Pricilla Philia Br. Purba",
					},
					{
						Nama:      "Akiyo Ramadhan Sarsito",
						Fakultas:  "Fakultas Hukum",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609802/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00411_jmegbj.jpg",
						Instagram: "kiyoooo__",
						LinkedIn:  "Akiyo Ramadhan Sarsito",
					},
				},
				Volunteer: []string{
					"Fadila Vairuz Harsyad",
					"Graciella",
					"Ivan Effendy",
					"Nisrina, Raniah",
				},
			},
			{
				ID:     4,
				Divisi: "Communication Editorial Marketing",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Maria Desvita Sari",
						Fakultas:  "Fakultas Ilmu Sosial dan Ilmu Politik",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609802/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00363_kvghlx.jpg",
						Instagram: "mareeavs",
						LinkedIn:  "Maria Desvita Sari",
					},
					{
						Nama:      "Tsani Adam Zidan Abidin",
						Fakultas:  "Fakultas Perikanan dan Ilmu Kelautan",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609801/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00386_gbue7q.jpg",
						Instagram: "adamzzdd",
						LinkedIn:  "adam-zidan",
					},
					{
						Nama:      "Muhammad Ja’far Shiddiq",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609801/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00378_du4rfx.jpg",
						Instagram: "jafarrs",
						LinkedIn:  "Muhammad Ja’far Shiddiq",
					},
					{
						Nama:      "Kania Ayu Ramadhina",
						Fakultas:  "Fakultas Ilmu Administrasi",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609812/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00372_pgrlls.jpg",
						Instagram: "kaniiaar",
						LinkedIn:  "kaniaayuramadhina",
					},
				},
				Volunteer: []string{
					"Muhammad Satria Bintang Pratama",
					"Wahyu Ramadhan Hamid Kuna",
					"Maureen Anabelle",
					"Amar Dzakwan",
				},
			},
			{
				ID:     5,
				Divisi: "Event Manager",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Amira Nada Fauziyyah",
						Fakultas:  "Fakultas Ilmu Sosial dan Ilmu Politik",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609809/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00341_z4dccv.jpg",
						Instagram: "amiranadaa",
						LinkedIn:  "Amira Nada Fauziyyah",
					},
					{
						Nama:      "Aloysius Ardeyno",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609810/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00336_xr796l.jpg",
						Instagram: "aloysiusardeyno",
						LinkedIn:  "aloysiusardeyno",
					},
					{
						Nama:      "Mastari Ardra Athaya",
						Fakultas:  "Fakultas Ilmu Sosial dan Ilmu Politik",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609810/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00347_dwiwxw.jpg",
						Instagram: "ardraath",
						LinkedIn:  "Mastari Ardra Athaya",
					},
				},
				Volunteer: []string{
					"Zaky Radja Nugroho",
					"Fadlan Fayudhi",
					"Calvin Jaya Septian ",
					"Nicolas Duarte",
					"Kayla Kamila Suprapto Putri",
				},
			},
			{
				ID:     6,
				Divisi: "Executive Producer",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Muhammad Atallah Azayaka",
						Fakultas:  "Fakultas Ilmu Sosial dan Ilmu Politik",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609799/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00740_fvogje.jpg",
						Instagram: "akaa_47",
						LinkedIn:  "Muhammad Atallah Azayaka",
					},
					{
						Nama:      "Andaru Anindito",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609812/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00743_yhylcd.jpg",
						Instagram: "daruanindito",
						LinkedIn:  "andaruanindito",
					},
				},
				Volunteer: []string{
					"Ibrahim Hussein",
					"RA Gantari Koos",
					"Nadila Syarifa Boediono",
					"Laisya Nauva",
				},
			},
			{
				ID:     7,
				Divisi: "Design",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Muhammad Faris Faisal",
						Fakultas:  "Fakultas Perikanan dan Ilmu Kelautan",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694845920/TEDXUB2023%20Our%20Team/Tambahan/Isal_otcjrq.png",
						Instagram: "_mffaisal",
						LinkedIn:  "mfarisfaisal",
					},
					{
						Nama:      "Alifia Halida Zahra",
						Fakultas:  "Fakultas Vokasi",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609810/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00318_hzn5nf.jpg",
						Instagram: "alifiahz",
						LinkedIn:  "Alifia Halida Zahra",
					},
					{
						Nama:      "Dafina Amira",
						Fakultas:  "Fakultas Teknik",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609811/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00303_xtyive.jpg",
						Instagram: "eyes0fdapp",
						LinkedIn:  "dafinamiraa",
					},
				},
				Volunteer: []string{
					"Najwa Nabela Fildza",
					"Achmad Zaini Eka A. S. ",
					"Ajeng Putri Ayu Winaryati",
				},
			},
			{
				ID:     8,
				Divisi: "Video Production",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Muhammad Arsa Alamsyah",
						Fakultas:  "Fakultas Ilmu Budaya",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609798/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00795_loxuef.jpg",
						Instagram: "soulfulwhisper____",
						LinkedIn:  "Muhammad Arsa Alamsyah",
					},
					{
						Nama:      "Febryan Kusuma Irawan",
						Fakultas:  "Fakultas Vokasi",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609799/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00785_l20zig.jpg",
						Instagram: "febryanirawan_",
						LinkedIn:  "Febryan Kusuma Irawan",
					},
					{
						Nama:      "Dio Rama Mahendra",
						Fakultas:  "Fakultas Hukum",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609813/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00354_xu3ez0.jpg",
						Instagram: "themahendra_",
						LinkedIn:  "Dio Rama Mahendra",
					},
				},
				Volunteer: []string{
					"Laura Ayako Suryawarman",
					"Muhammad Ariq Farhan",
					"Duncan",
				},
			},
			{
				ID:     9,
				Divisi: "Website Manager",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Galih Permana",
						Fakultas:  "Fakultas Ilmu Komputer",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609801/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00725_gk7vo9.jpg",
						Instagram: "galjhpermana",
						LinkedIn:  "galihpermana",
					},
					{
						Nama:      "Zidane Ali Basalamah",
						Fakultas:  "Fakultas Ilmu Komputer",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609799/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00707_ciehok.jpg",
						Instagram: "zidanebasalamah",
						LinkedIn:  "zidanebasalamah",
					},
					{
						Nama:      "Ryo Shandy",
						Fakultas:  "Fakultas Ilmu Komputer",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694941123/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00723_1-min_xt1ird.jpg",
						Instagram: "ryoshandy",
						LinkedIn:  "ryoshandy",
					},
				},
				Volunteer: []string{
					"Gede Indra Adi Brata",
					"Muhammad Farrel Reginaldo Ahnaf",
					"Regita Nadia Putri",
					"Leonard Eikel Arapenta Tarigan",
				},
			},
			{
				ID:     10,
				Divisi: "Sponsorship",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Ana Pralisti",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694845919/TEDXUB2023%20Our%20Team/Tambahan/Ana_tqiqzv.png",
						Instagram: "anapralisti",
						LinkedIn:  "Ana Pralisti",
					},
					{
						Nama:      "Silviana Rahel Novianty Hutasoit ",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694845919/TEDXUB2023%20Our%20Team/Tambahan/Rahel_z9gp3u.png",
						Instagram: "rahelnh",
						LinkedIn:  "silvianarahel",
					},
					{
						Nama:      "Dayang Puspa Sandya Faiza",
						Fakultas:  "Fakultas Matematika dan Ilmu Pengetahuan Alam",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609812/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00326_ap9eiu.jpg",
						Instagram: "dayangpuspaa_",
						LinkedIn:  "Dayang Puspa Sandya Faiza",
					},
				},
				Volunteer: []string{
					"Syahdad Nabil Mudzaffar",
					"Kishi Aisha Zafira",
					"Alfath Mar'ie Baihaqi",
				},
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
