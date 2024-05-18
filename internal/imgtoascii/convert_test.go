package imgtoascii

import (
	"log"
	"os"
	"testing"

	"github.com/pixiv/go-libjpeg/jpeg"
)

func BenchmarkConvert(b *testing.B) {
	pwd, _ := os.Getwd()
	// Photo by <a href="https://unsplash.com/pt-br/@madhatterzone?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Manja Vitolic</a> na <a href="https://unsplash.com/pt-br/fotografias/gato-preto-e-branco-deitado-na-cadeira-de-bambu-marrom-dentro-do-quarto-gKXKBY-C-Dk?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>
	imgPath := pwd + "/test.jpg"
	file, err := os.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file, &jpeg.DecoderOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		Convert(img, Options{})
	}
}
