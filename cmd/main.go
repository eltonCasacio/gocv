package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eltoncasacio/gocv/dto"
	"github.com/eltoncasacio/gocv/internal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gocv.io/x/gocv"
)

func main() {
	var filterNames []string

	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	showIMG := gocv.NewWindow("Detect")
	defer showIMG.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/border", func(w http.ResponseWriter, r *http.Request) {
		var params dto.BorderFiltersParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(params)
		filterNames = params.Filters
	})

	go http.ListenAndServe(":3000", r)

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", 0)
			return
		}
		if img.Empty() {
			continue
		}

		if img.Empty() {
			continue
		}

		internal.ApplyFilterr(&img, filterNames)

		showIMG.IMShow(img)
		if showIMG.WaitKey(1) >= 0 {
			break
		}
	}
}
