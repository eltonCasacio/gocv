package extracaocaracteristicas

import (
	"fmt"

	"gocv.io/x/gocv"
)

func Run(filename string) {
	fmt.Println(filename)
	img := gocv.IMRead(filename, gocv.IMReadGrayScale)
	window := gocv.NewWindow("Original")
	window2 := gocv.NewWindow("result")

	mat := gocv.NewMat()
	gocv.Threshold(img, &mat, float32(200), float32(255), gocv.ThresholdBinaryInv)

	contornos := gocv.FindContours(mat, gocv.RetrievalTree, gocv.ChainApproxSimple)
	area := gocv.ContourArea(contornos.At(0))
	fmt.Println("area: ", area)

	perimetro := gocv.ArcLength(contornos.At(0), true)
	fmt.Println("perimetro: ", perimetro)

	for {
		window.IMShow(img)
		window2.IMShow(mat)
		if window.WaitKey(10) >= 0 {
			break
		}
	}

}
