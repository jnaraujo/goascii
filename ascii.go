package goascii

import (
	"image"
	"image/color"
	"math"
	"strings"

	"github.com/disintegration/imaging"
)

type Converter struct {
	Chars []byte
}

func New() *Converter {
	// chars from https://paulbourke.net/dataformats/asciiart/
	var charlist = " .'`^\",:;Il!i><~+_-?][}{1)(|/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	var chars = []byte(charlist)

	return &Converter{
		Chars: chars,
	}
}

type Options struct {
	Columns int
	Filter  *imaging.ResampleFilter
}

var defaultOptions = &Options{
	Filter:  &imaging.Lanczos,
	Columns: 80,
}

func (c *Converter) Convert(img image.Image, options ...Options) (string, error) {
	option := mergeOptions(defaultOptions, options...)

	cols := img.Bounds().Dx()
	if option.Columns != 0 {
		cols = option.Columns
	}
	rows := int(float64(img.Bounds().Dy())*(float64(cols)/float64(img.Bounds().Dx()))) / 2

	var data strings.Builder
	data.Grow(cols * rows)

	if cols != img.Bounds().Dx() && rows != img.Bounds().Dy() {
		img = imaging.Resize(img, cols, rows, *option.Filter)
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			clr := img.At(x, y)
			gray := colorToGrayScale(clr)
			char := c.grayToChar(gray)
			data.WriteByte(char)
		}
		data.WriteByte('\n')
	}
	return data.String(), nil
}

func mergeOptions(defaultOptions *Options, options ...Options) *Options {
	if len(options) == 0 {
		return defaultOptions
	}

	option := options[0]
	if option.Filter == nil {
		option.Filter = defaultOptions.Filter
	}

	return &option
}

func colorToGrayScale(clr color.Color) uint8 {
	r, g, b, _ := clr.RGBA()
	return uint8((2126*r + 7152*g + 722*b) / 65536 / 256)
}

func (c *Converter) grayToChar(gray uint8) byte {
	var charLength = uint8(math.Ceil(256 / float64(len(c.Chars))))
	return c.Chars[gray/charLength]
}
