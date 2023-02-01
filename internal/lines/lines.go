// What it does:
//
// This example shows how to find lines in an image using Hough transform.
//
// How to run:
//
// 		go run ./cmd/find-lines/main.go lines.jpg
//

package lines

import (
	"image"
	"image/color"
	"math"

	"gocv.io/x/gocv"
)

type lines struct {
	img *gocv.Mat
}

func NewLines(img *gocv.Mat) *lines {
	return &lines{img: img}
}

func (l *lines) FindLines() {
	matCanny := gocv.NewMat()
	matLines := gocv.NewMat()

	gocv.Canny(*l.img, &matCanny, 100, 200)
	gocv.HoughLinesP(matCanny, &matLines, 1, math.Pi/180, 80)

	for i := 0; i < matLines.Rows(); i++ {
		pt1 := image.Pt(int(matLines.GetVeciAt(i, 0)[0]), int(matLines.GetVeciAt(i, 0)[1]))
		pt2 := image.Pt(int(matLines.GetVeciAt(i, 0)[2]), int(matLines.GetVeciAt(i, 0)[3]))
		gocv.Line(l.img, pt1, pt2, color.RGBA{0, 255, 20, 50}, 4)
	}
}
