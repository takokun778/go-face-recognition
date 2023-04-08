package main

import (
	"fmt"
	"go-face-recognition/internal"
)

func main() {
	target := "miu-chan.jpg"

	face, err := internal.ExtractFaceFeature(target)
	if err != nil {
		panic(err)
	}

	b, err := internal.FaceFeatureToBase64(face)
	if err != nil {
		panic(err)
	}

	fmt.Println(b)
}
