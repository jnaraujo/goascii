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

	// f, err := os.Create(pwd + "/cpu.prof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

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

	// for i := 0; i < 100; i++ {
	// 	file.Seek(0, 0)
	// 	_, err = imgtoascii.Convert(file)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	os.WriteFile(pwd+"/test/cat.txt", []byte(ascii), 0644)
}
