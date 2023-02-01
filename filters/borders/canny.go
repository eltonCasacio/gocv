/*
Imagem a ser trata em tom de cinza
2,3  Intensidade (quanto maior o valor, menos bordas detectadas)
*/
package borders

import (
	"gocv.io/x/gocv"
)

func Canny(filename string) {
	img := gocv.IMRead(filename, gocv.IMReadAnyColor)
	window := gocv.NewWindow("Canny - Imagem Original")
	window2 := gocv.NewWindow("Canny")
	defer window.Close()
	defer window2.Close()

	cannyIMG := gocv.NewMat()
	gocv.Canny(img, &cannyIMG, float32(100), float32(200))

	for {
		window.IMShow(img)
		window2.IMShow(cannyIMG)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
