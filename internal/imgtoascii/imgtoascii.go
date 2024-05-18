package imgtoascii

import (
	"image/color"
	"io"
	"strings"

	"github.com/pixiv/go-libjpeg/jpeg"
)

func Convert(file io.Reader) (string, error) {
	img, err := jpeg.Decode(file, &jpeg.DecoderOptions{})
	if err != nil {
		return "", err
	}

	var ascii strings.Builder
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			ascii.WriteString(grayToChar(colorToGrayScale(img.At(x, y))))
		}
		ascii.WriteString("\n")
	}
	return ascii.String(), nil
}

func colorToGrayScale(clr color.Color) uint8 {
	r, g, b, _ := clr.RGBA()
	return uint8((2126*r + 7152*g + 722*b) / 65536 / 256)
}

var chars = [26]string{"M", "N", "H", "#", "Q", "U", "A", "D", "O", "Y", "2", "6", "8", "Z", "0", "L", "C", "J", "P", "G", "9", "S", "I", ":", "!", " "}

func grayToChar(gray uint8) string {
	return chars[gray/10]
}
