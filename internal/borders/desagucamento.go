package borders

import (
	"image"

	"gocv.io/x/gocv"
)

func Desagucamento(src *gocv.Mat) *gocv.Mat {
	ressonancia_suavizada := gocv.NewMat()
	ressonancia_detalhada := gocv.NewMat()
	ressonancia_realce := gocv.NewMat()

	gocv.Add(ressonancia_detalhada, ressonancia_detalhada, &ressonancia_detalhada)
	gocv.Add(ressonancia_detalhada, ressonancia_detalhada, &ressonancia_detalhada)
	gocv.Add(ressonancia_detalhada, ressonancia_detalhada, &ressonancia_detalhada)
	gocv.Add(ressonancia_detalhada, ressonancia_detalhada, &ressonancia_detalhada)

	gocv.GaussianBlur(*src, &ressonancia_suavizada, image.Pt(7, 7), float64(0), float64(0), gocv.BorderDefault)
	gocv.Subtract(*src, ressonancia_suavizada, &ressonancia_detalhada)
	gocv.Add(*src, ressonancia_detalhada, &ressonancia_realce)
	return &ressonancia_realce
}
