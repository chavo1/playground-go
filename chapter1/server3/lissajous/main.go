package main

import (
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

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles := 5   // default number
	res := 0.001  // default
	size := 100   // default
	nframes := 64 // default
	delay := 8    // default

	for param, values := range r.Form {
		value := values[0]

		switch param {
		case "cycles":
			if i, err := strconv.Atoi(value); err == nil {
				cycles = i
			}
		}
	}

	lissajous(w, cycles, res, size, nframes, delay)
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(2*cycles)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // encoding errors ignore
}
