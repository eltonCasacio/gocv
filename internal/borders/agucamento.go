package borders

import (
	"gocv.io/x/gocv"
)

func Agucamento(src *gocv.Mat) {
	mat := src.Clone()
	gocv.Laplacian(mat, src, gocv.MatTypeCV8U, int(1), float64(1), float64(1), gocv.BorderDefault)
	gocv.Subtract(mat, *src, src)
}
