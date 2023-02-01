package borders

import (
	"gocv.io/x/gocv"
)

/*
	Operador de Sobel, utilizado para realçar contornos em imagens.
	Realça linhas verticais e horizontais mais escuras que o fundo sem realçar pontos isolados.
	Realçar arestas independente da direcao.
	As regioes destacadas por esse procedimento resulta em boras mais grossas comparado
	ao resultado obtido comparado a outras tecnicas.

	>>>	Mais utilizado para deteccao de boras. <<<

	O filtro de sobel pode ser aplicado como um metodo sobel da biblioteca GOCV
	requer os seguintes parametros obrigatórios:
		src gocv.Mat
		dst *gocv.Mat
		ddepth gocv.MatType
		dx int
		dy int
		ksize int
		scale float64
		delta float64
		borderType gocv.BorderType

	gocv.Sobel => https://docs.opencv.org/4.x/d4/d86/group__imgproc__filter.html#gacea54f142e81b6758cb6f375ce782c8d
*/

func Sobel(
	src *gocv.Mat,
	ddepth gocv.MatType,
	dx_H, dx_V, dy_H, dy_V, ksize_H, ksize_V int,
	scale, delta float64,
	borderType gocv.BorderType,
) error {
	sobel_horizontal := src.Clone()
	sobel_vertical := src.Clone()

	gocv.Sobel(sobel_horizontal, src, ddepth, dx_H, dy_H, ksize_H, scale, delta, borderType)
	gocv.Sobel(sobel_vertical, src, ddepth, dx_V, dy_V, ksize_V, scale, delta, borderType)

	return nil
}
