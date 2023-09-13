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
				ID:     11,
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
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609799/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00723_ldnf7u.jpg",
						Instagram: "ryoshandy",
						LinkedIn:  "ryoshandy",
					},
				},
				Volunteer: []string{
					"Gede Indra Adi Brata",
					"Muhammad Farrel Reginaldo Ahnaf",
					"Regita Nadia Putri",
					"Leonard, Eikel Arapenta Tarigan",
				},
			},
			{
				ID:     9,
				Divisi: "Designer",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Muhammad Faris Faisal",
						Fakultas:  "Fakultas Perikanan dan Ilmu Kelautan",
						ImageURL:  "png",
						Instagram: "_mffaisal",
						LinkedIn:  "mfarisfaisal",
					},
					{
						Nama:      "Alifia Halida Zahra",
						Fakultas:  "Fakultas Vokasi",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609810/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00318_hzn5nf.jpg",
						Instagram: "alifiahz",
						LinkedIn:  "",
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
					"Ajeng P,utri Ayu Winaryati",
				},
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
						LinkedIn:  "pricilla-philia-479449236",
					},
					{
						Nama:      "Akiyo Ramadhan Sarsito",
						Fakultas:  "Fakultas Hukum",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609802/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00411_jmegbj.jpg",
						Instagram: "kiyoooo__",
						LinkedIn:  "akiyo-ramadhan-sarsito-44b504225",
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
				ID:     5,
				Divisi: "Sponsorship",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Ana Pralisti",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "png",
						Instagram: "anapralisti",
						LinkedIn:  "ana-pralisti-969315235",
					},
					{
						Nama:      "Silviana Rahel Novianty Hutasoit ",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "png",
						Instagram: "rahelnh",
						LinkedIn:  "silvianarahel",
					},
					{
						Nama:      "Dayang Puspa Sandya Faiza",
						Fakultas:  "Fakultas Matematika dan Ilmu Pengetahuan Alam",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609812/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00326_ap9eiu.jpg",
						Instagram: "dayangpuspaa_",
						LinkedIn:  "dayang-puspa-sandya-faiza-186305174",
					},
				},
				Volunteer: []string{
					"Syahdad Nabil Mudzaffar",
					"Kishi Aisha Zafira",
					"Alfath ,Mar'ie Baihaqi",
				},
			},
			{
				ID:     10,
				Divisi: "Video Production",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Muhammad Arsa Alamsyah",
						Fakultas:  "Fakultas Ilmu Budaya",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609798/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00795_loxuef.jpg",
						Instagram: "soulfulwhisper____",
						LinkedIn:  "muhammad-arsa-alamsyah-9694ba228",
					},
					{
						Nama:      "Febryan Kusuma Irawan",
						Fakultas:  "Fakultas Vokasi",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609799/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00785_l20zig.jpg",
						Instagram: "febryanirawan_",
						LinkedIn:  "febryan-kusuma-irawan-b3bb6b233",
					},
					{
						Nama:      "Dio Rama Mahendra",
						Fakultas:  "Fakultas Hukum",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609813/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00354_xu3ez0.jpg",
						Instagram: "themahendra_",
						LinkedIn:  "dio-rama-mahendra-313a46283",
					},
				},
				Volunteer: []string{
					"Laura Ayako Suryawarman",
					"Muhammad Ariq Farhan",
					"Duncan",
				},
			},
			{
				ID:     6,
				Divisi: "Event Manager",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Amira Nada Fauziyyah",
						Fakultas:  "Fakultas Ilmu Sosial dan Ilmu Politik",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609809/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00341_z4dccv.jpg",
						Instagram: "amiranadaa",
						LinkedIn:  "amira-nada-fauziyyah-a42839198",
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
						LinkedIn:  "mastari-ardra-athaya-1b6b06253",
					},
				},
				Volunteer: []string{
					"Zaky Radja Nugroho",
					"Fadlan Fayudhi",
					"Calvin Jaya Septian ",
					"Nicolas Duarte",
					"Kayla K,amila Suprapto Putri",
				},
			},
			{
				ID:     7,
				Divisi: "Executive Producer",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Muhammad Atallah Azayaka",
						Fakultas:  "Fakultas Ilmu Sosial dan Ilmu Politik",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609799/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00740_fvogje.jpg",
						Instagram: "akaa_47",
						LinkedIn:  "",
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
					"Laisya ,Nauva",
				},
			},
			{
				ID:     8,
				Divisi: "Communication, Editorial, dan Marketing",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Maria Desvita Sari",
						Fakultas:  "Fakultas Ilmu Sosial dan Ilmu Politik",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609802/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00363_kvghlx.jpg",
						Instagram: "mareeavs",
						LinkedIn:  "maria-desvita-sari-0b2719237",
					},
					{
						Nama:      "Tsani Adam Zidan Abidin",
						Fakultas:  "Fakultas Perikanan dan Ilmu Kelautan",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609801/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00386_gbue7q.jpg",
						Instagram: "adamzzdd",
						LinkedIn:  "adam-zidan",
					},
					{
						Nama:      "Kristian Ja'far Manullang",
						Fakultas:  "Fakultas Kriminal",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694624146/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/jaffarv2_rjo3jt.jpg",
						Instagram: "husein_hadar",
						LinkedIn:  "",
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
				ID:     1,
				Divisi: "Orginizer",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Nabyl Fadhlurrahman",
						Fakultas:  "Fakultas Ilmu Sosial dan Ilmu Politik",
						ImageURL:  "",
						Instagram: "nabylf",
						LinkedIn:  "maria-desvita-sari-0b2719237",
					},
				},
				Volunteer: []string{},
			},
			{
				ID:     2,
				Divisi: "Co-Orginizer",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Muhammad Nur Mi'raj AL-Qadrie",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "",
						Instagram: "qadrieeee",
						LinkedIn:  "maria-desvita-sari-0b2719237",
					},
				},
				Volunteer: []string{},
			},
			{
				ID:     4,
				Divisi: "Communication, Editorial, dan Marketing",
				CoreTeam: []CoreTeam{
					{
						Nama:      "Reihansyah Ardhiya Amanullah",
						Fakultas:  "Fakultas Ekonomi dan Bisnis",
						ImageURL:  "https://res.cloudinary.com/dpcwbnax4/image/upload/v1694609800/TEDXUB2023%20Our%20Team/Website%20Low%20Res%20Rev/DSC00430_uufyrv.jpg",
						Instagram: "reihardhiya",
						LinkedIn:  "maria-desvita-sari-0b2719237",
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
