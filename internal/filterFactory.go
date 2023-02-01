package internal

import (
	"fmt"

	"github.com/eltoncasacio/gocv/internal/borders"

	"gocv.io/x/gocv"
)

//  int(1), int(0), int(1), float64(1), float64(1), gocv.BorderDefault)
// gocv.Sobel(*src, &sobel_vertical, gocv.MatTypeCV8U, int(0), int(1), int(1), float64(1), float64(1), gocv.BorderDefault)

func ApplyFilterr(src *gocv.Mat, filters []string) {
	for _, filter := range filters {
		switch filter {
		case "AGUCAMENTO":
			borders.Agucamento(src)
		case "CANNY":
			borders.Canny(src, 100, 200)
		case "SOBEL":
			err := borders.Sobel(src, gocv.MatTypeCV8U, 1, 1, 0, 0, 1, 1, float64(1), float64(1), gocv.BorderDefault)
			if err != nil {
				fmt.Println("Erro nao sei", err)
				continue
			}
			// case "Desagucamento":
			// 	borders.Desagucamento(src, dist)
			// case "Circle":
			// 	borders.Desagucamento(src, dist)
			// case "SOBEL":
			// 	borders.Desagucamento(src, dist)
		}
	}
}
