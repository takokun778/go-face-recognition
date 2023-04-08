package main

import (
	"fmt"
	"go-face-recognition/internal"
	"log"
	"path/filepath"

	"github.com/Kagami/go-face"
)

const (
	originBase64 = "eyJSZWN0YW5nbGUiOnsiTWluIjp7IlgiOjU2LCJZIjoxOTl9LCJNYXgiOnsiWCI6Njk4LCJZIjo4NDF9fSwiRGVzY3JpcHRvciI6Wy0wLjA3NDU3NDc3NiwwLjA0NzE2Nzk3LDAuMDQ0NjgyNTMzLC0wLjA4Mjg1MTIzLC0wLjE0NDc3ODEsLTAuMDU1NTc1MDk1LC0wLjA5Mjk3NjE2LC0wLjAzODQ3OTg3MiwwLjA5ODkxNTYyLC0wLjEwOTM3Mjc5NSwwLjE1ODc3NTcyLC0wLjEwMTk0Njk3LC0wLjE5NzI1OTk4LC0wLjAyOTc2NTYxLC0wLjAxOTkzNDY5NywwLjIxNzk5MTI2LC0wLjE3NDY4MzY1LC0wLjE5MDAzMTg0LC0wLjAzMzY5MDkxNSwwLjA0MTA5MjQ3OCwwLjA0ODk4NDE2NiwwLjAzODk5NzI2LC0wLjAzMjYxODUzNCwwLjA3NTc5Njg3LC0wLjE2ODI5MDY0LC0wLjM0NTQzNjcsLTAuMDM2OTkxLC0wLjA2NjE1NzM3LC0wLjA2MjIxODMyNywtMC4wMjc1MDU5ODMsLTAuMDM1MzM2MTkzLDAuMDY1ODE4OTMsLTAuMTU4NjM4MDksLTAuMDAwODk2MTk4NywwLjA5MzgxOTc4LDAuMTA1ODE0NjY2LDAuMDI1MjEyNTczLC0wLjA2OTMxMjMxLDAuMTczNjE0NjgsLTAuMDE5NDE3NDgzLC0wLjM2NDQ4ODcsMC4wNTMzNzA3NDQsMC4wOTAxMjkyNTYsMC4yMjYwMDAyNiwwLjE2NTQwOTIyLDAuMDIzNDA5NzI0LDAuMDA3MjczOTQ4LC0wLjEyNjk5Mzk1LDAuMTA3ODE2OTUsLTAuMTk4OTE5MDksMC4wMDc5MDE5MjksMC4xNDQzNDI0OCwwLjAzMTQ4MDA1NSwwLjA0NjQ4MzMyNywtMC4wMDY3OTA0MTQ1LC0wLjE0NDg2NzgyLDAuMDQ0Mjk2ODksMC4xMTY2NjcxMiwtMC4xMzQxNDE2OCwtMC4wNDEzNTU3OTYsMC4wODg2NzE0NSwtMC4wNDE3ODMsMC4wMzQxNjQxMTIsLTAuMTQ0ODk4LDAuMjM0NDI4MiwwLjAxODc3ODE0MiwtMC4xNTY4NjQ2LC0wLjE3NzcyMDY0LDAuMDg4MDcwMDk1LC0wLjE3Njg1Mzg3LC0wLjExNjc2MTgyLDAuMDc1ODc5MDQ1LC0wLjEzMTI0ODAzLC0wLjE4NDYwMzc3LC0wLjM0MDM5NDQ0LC0wLjAzNTk2MjgyLDAuMzEwMzAwMTcsMC4xNjA0NjQ0NywtMC4yMjA0MTI5NywwLjA3ODY5MzE5LC0wLjAwNDE5NzIxMSwwLjA0MjY5MDY4NywwLjExMzI3MTQ3NSwwLjE4MTQ5NjIsMC4wMjA4MjgxOTMsMC4wODQ4NTA2OSwtMC4wNTQ3Mjk3MzQsLTAuMDIzMTU4MzM0LDAuMjc1NDUyNTIsLTAuMDQzMDA0MDc3LC0wLjAwMDk3NTI3MTgsMC4yMTUzNjg4OCwwLjAyNDE3MDg5NCwwLjA5MjMxMDYyLDAuMDA3ODAwMjE4LDAuMDA5NDQ2NTI3LC0wLjA0MDk0NDU1LDAuMDA2NzU1NDA2LC0wLjE0MTk3NjE2LDAuMDM1MzU1MDU4LDAuMDAyNzQ2MDg0NywwLjAyMzA1NDgxNCwwLjAwNzg0NDM5NSwwLjE1NzYxNDg5LC0wLjE4MDQzNTU1LDAuMTQzODgyMiwtMC4wMTI3NzUyNzIsMC4wMzQ3NjYyMSwtMC4wMDg3MTE0MjY1LDAuMTA0MzkwMzIsLTAuMDU0MDcxMzksLTAuMDgzODc1MzYsMC4xMTQ2ODM0OTQsLTAuMjEyODk0MTYsMC4xNzU0MzQ1NywwLjE5MjUzMDExLDAuMTA5Njk0NDM2LDAuMDg0NjQyMDE1LDAuMTUxODkxNTQsMC4wMzIzNjE3NjQsLTAuMDMyOTkyNTEyLC0wLjAwNzcxMzc2NCwtMC4yMzg0MTY2LC0wLjAyNzI5OTAxNSwwLjExMjI1ODY3LC0wLjA3NzA3NDYyLDAuMTI3MzYzNiwwLjA0MTQwNzY3XSwiU2hhcGVzIjpbeyJYIjo1NjEsIlkiOjM4OH0seyJYIjo0NDAsIlkiOjM5Nn0seyJYIjoxNzksIlkiOjM3MH0seyJYIjoyODEsIlkiOjM4N30seyJYIjozNDAsIlkiOjYzMX1dfQ=="
)

func main() {
	target := "target.jpg"

	if RecognizeTarget(target) {
		fmt.Println("Recognized")
	} else {
		fmt.Println("Not Recognized")
	}
}

func RecognizeTarget(targetImg string) bool {
	rec, err := face.NewRecognizer(internal.DataDir)
	if err != nil {
		fmt.Println("Cannot initialize recognizer")
	}
	defer rec.Close()

	faces, err := rec.RecognizeFile(filepath.Join(internal.DataDir, targetImg))
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

	miuChan, err := internal.Base64ToFaceFeature(originBase64)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if miuChan == nil {
		log.Fatalf("No faces found")
	}

	match := rec.ClassifyThreshold(miuChan.Descriptor, internal.Tolerance)

	return match >= 0
}
