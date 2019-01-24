package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	imagePath := flag.String("image", "", "path to image")
	flag.Parse()

	if *imagePath == "" {
		log.Fatal("Please specify image path. Use --help for help.")
	}

	//read image and return reader
	// read the image file
	m, err := os.Open(*imagePath)
	if err != nil {
		log.Fatalln("Error opening image", err)
	}
	defer m.Close()

	start := time.Now()
	emotion, err := getEmotion(m)
	duration := time.Since(start)

	fmt.Printf("Here is the full JSON response:%s\n", emotion)
	fmt.Printf("Time to call Face API: %s\n", duration)
}
