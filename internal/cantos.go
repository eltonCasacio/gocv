package internal

import (
	"gocv.io/x/gocv"
)

func Cantos(filename string) {
	img := gocv.IMRead(filename, gocv.IMReadGrayScale)

	window := gocv.NewWindow("Imagem Original")
	window2 := gocv.NewWindow("Cantos")
	defer window.Close()

	cantos := gocv.NewMat()

	gocv.GoodFeaturesToTrack(img, &cantos, int(4), float64(1), float64(2))

	for {
		window.IMShow(img)
		window2.IMShow(cantos)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
