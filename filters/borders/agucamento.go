package borders

import (
	"gocv.io/x/gocv"
)

func Agucamento(filename string) {
	window := gocv.NewWindow("Ressonancia - original")
	window_laplacian := gocv.NewWindow("ressonancia - laplacian")

	original_img := gocv.IMRead(filename, gocv.IMReadGrayScale)

	laplacian := gocv.NewMat()
	gocv.Laplacian(original_img, &laplacian, gocv.MatTypeCV8U, int(1), float64(1), float64(1), gocv.BorderDefault)
	gocv.Subtract(original_img, laplacian, &laplacian)
	// gocv.AddWeighted(laplacian, float64(1), laplacian, float64(1), float64(1), &laplacian)

	for {
		window.IMShow(original_img)
		window_laplacian.IMShow(laplacian)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
