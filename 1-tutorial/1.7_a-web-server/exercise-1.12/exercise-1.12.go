// server4 prints a lissajous GIF out the web browser.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.RGBA{0x21, 0xa1, 0x21, 0xff}}

const (
	whiteIndex = 0 // first color in pallete
	greenIndex = 1 // second color in palette
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// renders a lissajous GIF
func handler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	lissajous(w, req)
}

func lissajous(output io.Writer, req *http.Request) {
	var (
		cycles  float64 = 5     // number of complete x oscillator revolutions
		res             = 0.001 // angular resolution
		size    int     = 100   // image canvas covers [-size...+size]
		nframes         = 64    // number of animation frames
		delay           = 8     // delay between frames in 10ms units
	)

	if v := req.Form.Get("cycles"); v != "" {
		cyclesForm, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintf(output, "Failed to parse the request form data...")
			return
		}
		cycles = float64(cyclesForm)
	}

	if v := req.Form.Get("delay"); v != "" {
		delayForm, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintf(output, "Failed to parse the request form data...")
			return
		}
		delay = delayForm
	}

	if v := req.Form.Get("nframes"); v != "" {
		nframesForm, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintf(output, "Failed to parse the request form data...")
			return
		}
		nframes = nframesForm
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(output, &anim) // ignoring encoding errors
}
