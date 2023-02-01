package borders

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

func Circle(filename string) {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tgo run . [imgfile]")
		return
	}

	window := gocv.NewWindow("detected colors")
	window2 := gocv.NewWindow("end result")
	window3 := gocv.NewWindow("lower bound")
	window4 := gocv.NewWindow("upper bound")
	window2.MoveWindow(600, 0)
	window3.MoveWindow(0, 600)
	window4.MoveWindow(600, 600)
	defer window.Close()
	defer window2.Close()
	defer window3.Close()
	defer window4.Close()

	img := gocv.IMRead(filename, gocv.IMReadColor)
	defer img.Close()

	original_image := img.Clone()
	defer original_image.Close()

	gocv.MedianBlur(img, &img, 3)

	hsv_img := gocv.NewMat()
	defer hsv_img.Close()

	// yellow := gocv.NewScalar(0, 255, 255, 0)
	// yellow_mat := gocv.NewMatFromScalar(yellow, gocv.MatTypeCV8UC3)
	// gocv.CvtColor(yellow_mat, &yellow_mat, gocv.ColorBGRToHSV)
	// hsv := gocv.Split(yellow_mat)
	// fmt.Printf("H: %d S: %d V: %d\n", hsv[0].GetUCharAt(0, 0), hsv[1].GetUCharAt(0, 0), hsv[2].GetUCharAt(0, 0))

	gocv.CvtColor(img, &hsv_img, gocv.ColorBGRToHSV)
	img_rows, img_cols := hsv_img.Rows(), hsv_img.Cols()

	// lb1 := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(0.0, 208.0, 94.0, 0.0), img_rows, img_cols, gocv.MatTypeCV8UC3)
	// ub1 := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(179.0, 255.0, 255.0, 0.0), img_rows, img_cols, gocv.MatTypeCV8UC3)
	lb1 := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(20.0, 50.0, 50.0, 0.0), img_rows, img_cols, gocv.MatTypeCV8UC3)
	ub1 := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(40.0, 255.0, 255.0, 0.0), img_rows, img_cols, gocv.MatTypeCV8UC3)

	lb2 := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(155.0, 100.0, 100.0, 0.0), img_rows, img_cols, gocv.MatTypeCV8UC3)
	ub2 := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(180.0, 255.0, 255.0, 0.0), img_rows, img_cols, gocv.MatTypeCV8UC3)

	lower_bound := gocv.NewMat()
	upper_bound := gocv.NewMat()
	color_isolated_img := gocv.NewMat()
	circles := gocv.NewMat()
	defer lower_bound.Close()
	defer upper_bound.Close()
	defer color_isolated_img.Close()
	defer circles.Close()

	gocv.InRange(hsv_img, lb1, ub1, &lower_bound)
	gocv.InRange(hsv_img, lb2, ub2, &upper_bound)

	gocv.AddWeighted(lower_bound, 1.0, upper_bound, 1.0, 0.0, &color_isolated_img)
	gocv.GaussianBlur(color_isolated_img, &color_isolated_img, image.Pt(9, 9), 2, 2, gocv.BorderDefault)

	gocv.HoughCirclesWithParams(
		color_isolated_img,
		&circles,
		gocv.HoughGradient,
		1,
		float64(color_isolated_img.Rows()/8),
		100,
		20,
		0,
		0,
	)

	green := color.RGBA{0, 255, 0, 0}

	for i := 0; i < circles.Cols(); i++ {
		v := circles.GetVecfAt(0, i)
		fmt.Println(v)
		if len(v) > 2 {
			x := int(v[0])
			y := int(v[1])
			r := int(v[2])

			gocv.Circle(&original_image, image.Pt(x, y), r, green, 2)
		}
	}

	for {
		// window.IMShow(color_isolated_img)
		window2.IMShow(original_image)
		window3.IMShow(lower_bound)
		window4.IMShow(upper_bound)

		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
