package morfologicas

import (
	"image"

	"gocv.io/x/gocv"
)

func Abertura(filename string) {
	window := gocv.NewWindow("imagem original")
	window2 := gocv.NewWindow("imagem abertura")
	img := gocv.IMRead(filename, gocv.IMReadGrayScale)
	abertura := gocv.NewMat()

	estruturante := gocv.GetStructuringElement(gocv.MorphCross, image.Pt(2, 2))
	gocv.MorphologyEx(img, &abertura, gocv.MorphOpen, estruturante)

	for {
		window.IMShow(img)
		window2.IMShow(abertura)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
