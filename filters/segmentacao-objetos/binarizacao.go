package segmentacaoobjetos

import "gocv.io/x/gocv"

func Binarizacao(filename string) {
	window := gocv.NewWindow("Ressonancia - original")
	window2 := gocv.NewWindow("ressonancia - laplacian")

	img := gocv.IMRead(filename, gocv.IMReadGrayScale)
	segmentada := gocv.NewMat()

	gocv.Threshold(img, &segmentada, 100, 255, gocv.ThresholdBinary)

	for {
		window.IMShow(img)
		window2.IMShow(segmentada)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
