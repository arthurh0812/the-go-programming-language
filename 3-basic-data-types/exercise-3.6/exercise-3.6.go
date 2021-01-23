// mandelbrot emits a PNG image of the mandelbrot fractal. (depixeled)
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		epsX                   = (xmax - xmin) / width
		epsY                   = (ymax - ymin) / height
	)

	offX := [2]float64{-epsX, epsX}
	offY := [2]float64{-epsY, epsY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// supersampling:
			subpixels := make([]color.Color, 4)
			for xi := 0; xi < 2; xi++ {
				for yi := 0; yi < 2; yi++ {
					z := complex(x+offX[xi], y+offY[yi])
					subpixels = append(subpixels, mandelbrot(z))
				}
			}
			img.Set(px, py, avg(subpixels))
		}
	}

	file, err := os.OpenFile("img.png", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file, img)

	file.Close()
}

func avg(colors []color.Color) color.Color {
	var avgR, avgG, avgB, avgA uint8
	n := len(colors)
	for _, c := range colors {
		clrR, clrG, clrB, clrA := c.RGBA()
		avgR += uint8(clrR / uint32(n))
		avgG += uint8(clrG / uint32(n))
		avgB += uint8(clrB / uint32(n))
		avgA += uint8(clrA / uint32(n))
	}

	return color.RGBA{R: avgR, G: avgG, B: avgB, A: avgA}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50: // dark red
				return color.RGBA{100, 0, 0, 255}
			default:
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{R: 0, G: 0, B: 255 - uint8(logScale*255), A: 255}
			}
		}
	}

	return color.RGBA{R: 0, G: 0, B: 0, A: 255}
}
