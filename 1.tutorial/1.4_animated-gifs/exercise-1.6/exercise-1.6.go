package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.White,
	color.RGBA{0x21, 0xa1, 0x21, 0xff},
	color.RGBA{0xe0, 0x18, 0x18, 0xff},
	color.RGBA{0xe0, 0x18, 0xb5, 0xff},
	color.RGBA{0x18, 0xe0, 0xc5, 0xff},
	color.RGBA{0xe0, 0xbb, 0x18, 0xff},
	color.RGBA{0xa0, 0x52, 0x2d, 0xff},
	color.RGBA{0xa0, 0x2d, 0x70, 0xff},
}

const (
	whiteIndex = 0 // first color in pallete
	greenIndex = 1 // second color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(output io.Writer) {
	const (
		cycles  = 5       // number of complete x oscillator revolutions
		res     = 0.00001 // angular resolution
		size    = 300     //image canvas covers [-size...+size]
		nframes = 64      // number of animation frames
		delay   = 6       // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		var randomColor uint8
		if i%4 == 0 {
			randomColor = uint8(rand.Intn(6))
		}
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), randomColor+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(output, &anim) // ignoring encoding errors
}
