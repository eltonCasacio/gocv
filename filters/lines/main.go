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
	"fmt"
	"image"
	"image/color"
	"math"

	"gocv.io/x/gocv"
)

func Lines(filename string) {
	mat := gocv.IMRead(filename, gocv.IMReadColor)

	matCanny := gocv.NewMat()
	matLines := gocv.NewMat()

	window := gocv.NewWindow("detected lines")

	gocv.Canny(mat, &matCanny, 100, 200)
	gocv.HoughLinesP(matCanny, &matLines, 1, math.Pi/180, 80)

	fmt.Println(matLines.Cols())
	fmt.Println(matLines.Rows())
	for i := 0; i < matLines.Rows(); i++ {
		pt1 := image.Pt(int(matLines.GetVeciAt(i, 0)[0]), int(matLines.GetVeciAt(i, 0)[1]))
		pt2 := image.Pt(int(matLines.GetVeciAt(i, 0)[2]), int(matLines.GetVeciAt(i, 0)[3]))
		gocv.Line(&mat, pt1, pt2, color.RGBA{0, 255, 20, 50}, 4)
	}

	for {
		window.IMShow(mat)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
