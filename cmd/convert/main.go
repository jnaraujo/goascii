package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jnaraujo/goascii"
	"github.com/pixiv/go-libjpeg/jpeg"
)

func main() {
	pwd, _ := os.Getwd()

	imgPath := pwd + "/test/test.jpg"

	file, err := os.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()
	img, err := jpeg.Decode(file, &jpeg.DecoderOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Since(start))

	converter := goascii.New()

	start = time.Now()
	ascii, err := converter.Convert(img, goascii.Options{
		Columns: 80,
	})
	fmt.Println(time.Since(start))
	if err != nil {
		log.Fatal(err)
	}

	txt := strings.Replace(ascii, "*", " ", -1)

	os.WriteFile(pwd+"/test/test.txt", []byte(txt), 0644)
}
