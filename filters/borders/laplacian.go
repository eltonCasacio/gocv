package borders

import (
	"image"

	"gocv.io/x/gocv"
)

// Operador Laplaciano, utilizado para realçar bordas - Filtro espacial de ordem 3
// gocv.Laplacian => https://docs.opencv.org/4.x/d4/d86/group__imgproc__filter.html#gad78703e4c8fe703d479c1860d76429e6

func Laplacian(filename string) {
	window := gocv.NewWindow("Garagem - original")
	window_laplacian := gocv.NewWindow("Garagem - laplacian")

	original_img := gocv.IMRead(filename, gocv.IMReadGrayScale)
	laplacian := gocv.NewMat()

	gocv.GaussianBlur(original_img, &laplacian, image.Pt(7, 7), float64(8), float64(8), gocv.BorderDefault)
	gocv.Add(laplacian, laplacian, &laplacian)
	gocv.Add(laplacian, laplacian, &laplacian)
	gocv.Laplacian(laplacian, &laplacian, gocv.MatTypeCV8U, int(1), float64(1), float64(1), gocv.BorderDefault)

	for {
		window.IMShow(original_img)
		window_laplacian.IMShow(laplacian)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
