package morfologicas

import (
	"image"

	"gocv.io/x/gocv"
)

func Erosao(filename string) {
	window := gocv.NewWindow("imagem original")
	window2 := gocv.NewWindow("imagem erosao")
	img := gocv.IMRead(filename, gocv.IMReadAnyColor)
	estruturante := gocv.GetStructuringElement(gocv.MorphCross, image.Pt(3, 3))

	erosao := gocv.NewMat()
	gocv.Erode(img, &erosao, estruturante)

	for {
		window.IMShow(img)
		window2.IMShow(erosao)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
