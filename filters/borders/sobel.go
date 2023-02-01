package borders

import (
	"gocv.io/x/gocv"
)

// Operador de Sobel, utilizado para realçar contornos em imagens
// gocv.Sobel => https://docs.opencv.org/4.x/d4/d86/group__imgproc__filter.html#gacea54f142e81b6758cb6f375ce782c8d

func Sobel(filename string) {
	img := gocv.IMRead(filename, gocv.IMReadGrayScale)
	sobel_horizontal := gocv.NewMat()
	sobel_vertical := gocv.NewMat()
	window := gocv.NewWindow("imagem original")
	whorizontal := gocv.NewWindow("horizontal")
	wvertical := gocv.NewWindow("vertical")

	gocv.Sobel(img, &sobel_horizontal, gocv.MatTypeCV8U, int(1), int(0), int(1), float64(1), float64(1), gocv.BorderDefault)
	gocv.Sobel(img, &sobel_vertical, gocv.MatTypeCV8U, int(0), int(1), int(1), float64(1), float64(1), gocv.BorderDefault)

	for {
		window.IMShow(img)
		whorizontal.IMShow(sobel_horizontal)
		wvertical.IMShow(sobel_vertical)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
