package borders

import (
	"image"

	"gocv.io/x/gocv"
)

// Operador Laplaciano, utilizado para realçar bordas - Filtro espacial de ordem 3
// gocv.Laplacian => https://docs.opencv.org/4.x/d4/d86/group__imgproc__filter.html#gad78703e4c8fe703d479c1860d76429e6

func Laplacian(src *gocv.Mat) {
	laplacian := gocv.NewMat()

	gocv.GaussianBlur(*src, &laplacian, image.Pt(7, 7), float64(8), float64(8), gocv.BorderDefault)
	gocv.Add(laplacian, laplacian, &laplacian)
	gocv.Add(laplacian, laplacian, &laplacian)
	gocv.Laplacian(laplacian, &laplacian, gocv.MatTypeCV8U, int(1), float64(1), float64(1), gocv.BorderDefault)
}
