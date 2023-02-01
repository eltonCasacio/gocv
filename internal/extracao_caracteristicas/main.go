package extracaocaracteristicas

import (
	"gocv.io/x/gocv"
)

type extractCaracteristic struct {
	img       *gocv.Mat
	contornos gocv.PointsVector
	area      float64
	perimetro float64
}

func NewExtractCaracteristic(img *gocv.Mat) *extractCaracteristic {
	return &extractCaracteristic{img: img}
}

func (e *extractCaracteristic) Extract() {
	mat := gocv.NewMat()
	gocv.Threshold(*e.img, &mat, float32(200), float32(255), gocv.ThresholdBinaryInv)

	e.contornos = gocv.FindContours(mat, gocv.RetrievalTree, gocv.ChainApproxSimple)
	e.area = gocv.ContourArea(e.contornos.At(0))
	e.perimetro = gocv.ArcLength(e.contornos.At(0), true)
}

func (e *extractCaracteristic) GetContorno() gocv.PointsVector {
	return e.contornos
}

func (e *extractCaracteristic) GetArea() float64 {
	return e.area
}

func (e *extractCaracteristic) GetPerimetro() float64 {
	return e.perimetro
}
