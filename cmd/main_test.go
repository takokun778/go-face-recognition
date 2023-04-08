package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Kagami/go-face"
)

const (
	dataDir = "images"
	// tolerance = 0.13
	tolerance = 0.13
)

func RecognizeTest() {
	for i := 1; i <= 9; i++ {
		Recognize(fmt.Sprintf("matsuri-%d.jpg", i))
	}
}

func Recognize(targetFile string) {
	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		fmt.Println("Cannot initialize recognizer")
	}
	defer rec.Close()

	faces, err := rec.RecognizeFile(filepath.Join(dataDir, targetFile))
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if len(faces) == 0 {
		log.Fatalf("No faces found")
	}

	var targets []face.Descriptor
	var indexs []int32

	for i, f := range faces {
		targets = append(targets, f.Descriptor)
		indexs = append(indexs, int32(i))
	}

	rec.SetSamples(targets, indexs)

	miuChan, err := rec.RecognizeSingleFile(filepath.Join(dataDir, "miu-1.jpg"))
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if miuChan == nil {
		log.Fatalf("No faces found")
	}

	match := rec.ClassifyThreshold(miuChan.Descriptor, tolerance)
	if match < 0 {
		log.Printf("No match found")

		return
	}

	log.Printf("Match found %d", match)
}
