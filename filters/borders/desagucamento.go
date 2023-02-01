package borders

import (
	"image"

	"gocv.io/x/gocv"
)

func Desagucamento(filename string) {
	window := gocv.NewWindow("Ressonancia - Original")
	window_ressonancia_suavizada := gocv.NewWindow("ressonancia_suavizada")
	window_ressonancia_detalhada := gocv.NewWindow("ressonancia_detalhada")
	window_ressonancia_realce := gocv.NewWindow("ressonancia_realce")

	img := gocv.IMRead(filename, gocv.IMReadGrayScale)

	ressonancia_suavizada := gocv.NewMat()
	ressonancia_detalhada := gocv.NewMat()
	ressonancia_realce := gocv.NewMat()

	gocv.Add(ressonancia_detalhada, ressonancia_detalhada, &ressonancia_detalhada)
	gocv.Add(ressonancia_detalhada, ressonancia_detalhada, &ressonancia_detalhada)
	gocv.Add(ressonancia_detalhada, ressonancia_detalhada, &ressonancia_detalhada)
	gocv.Add(ressonancia_detalhada, ressonancia_detalhada, &ressonancia_detalhada)

	gocv.GaussianBlur(img, &ressonancia_suavizada, image.Pt(7, 7), float64(0), float64(0), gocv.BorderDefault)
	gocv.Subtract(img, ressonancia_suavizada, &ressonancia_detalhada)
	gocv.Add(img, ressonancia_detalhada, &ressonancia_realce)

	for {
		window.IMShow(img)
		window_ressonancia_suavizada.IMShow(ressonancia_suavizada)
		window_ressonancia_detalhada.IMShow(ressonancia_detalhada)
		window_ressonancia_realce.IMShow(ressonancia_realce)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
