package detectedcolor

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gocv.io/x/gocv"
)

func Run(deviceID int) {

	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("start reading camera device: %v\n", deviceID)
	window := gocv.NewWindow("Detector de face humana")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	mat := gocv.NewMat()
	defer mat.Close()

	for {
		mat.ConvertTo(&mat, gocv.MatTypeCV8U)

		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		gocv.CvtColor(img, &img, gocv.ColorBGRToHSV)

		readFile, err := os.Open("conf.txt")
		if err != nil {
			panic(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var lines []string
		for fileScanner.Scan() {
			lines = append(lines, fileScanner.Text())
		}

		lb1, _ := strconv.ParseFloat(strings.Split(strings.Split(lines[0], "=")[1], ";")[0], 64)
		lb2, _ := strconv.ParseFloat(strings.Split(strings.Split(lines[1], "=")[1], ";")[0], 64)
		lb3, _ := strconv.ParseFloat(strings.Split(strings.Split(lines[2], "=")[1], ";")[0], 64)
		lb4, _ := strconv.ParseFloat(strings.Split(strings.Split(lines[3], "=")[1], ";")[0], 64)

		ub1, _ := strconv.ParseFloat(strings.Split(strings.Split(lines[4], "=")[1], ";")[0], 64)
		ub2, _ := strconv.ParseFloat(strings.Split(strings.Split(lines[5], "=")[1], ";")[0], 64)
		ub3, _ := strconv.ParseFloat(strings.Split(strings.Split(lines[6], "=")[1], ";")[0], 64)
		ub4, _ := strconv.ParseFloat(strings.Split(strings.Split(lines[7], "=")[1], ";")[0], 64)

		gocv.InRangeWithScalar(img,
			gocv.NewScalar(float64(lb1), float64(lb2), float64(lb3), float64(lb4)),
			gocv.NewScalar(float64(ub1), float64(ub2), float64(ub3), float64(ub4)), &mat)

		window.IMShow(mat)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
