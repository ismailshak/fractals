package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	rows       = 400
	columns    = 400
	minReal    = -2.0
	maxReal    = 1.0
	minIm      = -1.2
	maxIm      = minIm + (maxReal-minReal)*(rows/columns)
	realFactor = (maxReal - minReal) / (columns - 1)
	imFactor   = (maxIm - minIm) / (rows - 1)
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, columns, rows))
	log.Println("Starting render...")
	render(img)
	log.Println("Render complete.")
	log.Println("Encoding image...")
	f, err := os.Create("result.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	log.Println("Image is ready.")
}

/*
	z = z2 + c
	- breaking up complex numbers by their real and imaginary parts
*/
func render(img *image.RGBA) {

	for y := 0; y < rows; y++ {
		cIm := calculateComplexImaginary(y)
		for x := 0; x < columns; x++ {
			cReal := calculateComplexReal(x)

			zReal, zIm := cReal, cIm

			maxIterations := 30
			isInside := true

			for n := 0; n < maxIterations; n++ {
				zRealSquared, zImSquared := zReal*zReal, zIm*zIm

				if zRealSquared+zImSquared > 4 {
					isInside = false
					img.SetRGBA(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
					break
				}
				zIm = 2.0*zReal*zIm + cIm
				zReal = zRealSquared - zImSquared + cReal
			}
			if isInside {
				img.SetRGBA(x, y, color.RGBA{R: 0, G: 0, B: 0, A: 255})
			}
		}
	}

}

func calculateComplexReal(x int) float64 {
	return minReal + float64(x)*realFactor
}

func calculateComplexImaginary(y int) float64 {
	return minIm + float64(y)*imFactor
}
