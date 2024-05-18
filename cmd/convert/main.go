package main

import (
	"fmt"
	"img_to_ascii/internal/imgtoascii"
	"log"
	"os"
	"time"
)

func main() {
	pwd, _ := os.Getwd()

	imgPath := pwd + "/test/cat.jpg"

	file, err := os.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()
	ascii, err := imgtoascii.Convert(file)
	fmt.Println(time.Since(start))
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(pwd+"/test/cat.txt", []byte(ascii), 0644)
}
