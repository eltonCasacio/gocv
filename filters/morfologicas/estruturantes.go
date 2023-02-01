package morfologicas

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func Estruturante() {
	// rect := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(10, 10))
	// eliptico := gocv.GetStructuringElement(gocv.MorphEllipse, image.Pt(12, 12))
	cross := gocv.GetStructuringElement(gocv.MorphCross, image.Pt(12, 12))

	for l := 0; l < cross.Rows(); l++ {
		for c := 0; c < cross.Rows(); c++ {
			t := cross.GetSCharAt(l, c)
			fmt.Print(" ", t)
		}
		println()
	}
}
