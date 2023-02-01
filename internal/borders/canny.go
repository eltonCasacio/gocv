/*
Imagem a ser trata em tom de cinza
2,3  Intensidade (quanto maior o valor, menos bordas detectadas)
*/
package borders

import (
	"gocv.io/x/gocv"
)

func Canny(src *gocv.Mat, t1, t2 float32) error {
	mat := src.Clone()
	gocv.Canny(mat, src, t1, t2)
	return nil
}
