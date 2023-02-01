package main

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func Tensorflow(device int) {
	webcam, err := gocv.VideoCaptureDevice(device)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Detector de face humana")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	net := gocv.ReadNetFromTensorflow("/Users/eltoncasacio/testegocv/tensorflow/tensorflow_inception_graph.pb")

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", 0)
			return
		}
		if img.Empty() {
			continue
		}

		// convert to a 224x244 image blob that can be processed by Tensorflow
		blob := gocv.BlobFromImage(img, 1.0, image.Pt(224, 244), gocv.NewScalar(0, 0, 0, 0), true, false)
		defer blob.Close()

		// feed the blob into the classifier
		net.SetInput(blob, "input")

		// run a forward pass thru the network
		prob := net.Forward("softmax2")
		defer prob.Close()

		// reshape the results into a 1x1000 matrix
		probMat := prob.Reshape(1, 1)
		defer probMat.Close()

		// determine the most probable classification, and display it
		_, maxVal, _, maxLoc := gocv.MinMaxLoc(probMat)
		fmt.Printf("maxLoc: %v, maxVal: %v\n", maxLoc, maxVal)

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
