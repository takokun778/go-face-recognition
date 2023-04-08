package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Kagami/go-face"
)

const (
	dataDir   = "images"
	tolerance = 0.13
)

func main() {
	target := "miu-4.jpg"

	faces, err := ExtractFaceFeature(dataDir)
	if err != nil {
		panic(err)
	}

	bs := make([]string, 0, len(faces))
	for _, f := range faces {
		b, err := FaceFeatureToBase64(f)
		if err != nil {
			log.Printf("can't convert to base64: %v", err)

			continue
		}

		bs = append(bs, b)
	}

	const filePath = "faces.txt"

	if err := WriteFile(filePath, bs); err != nil {
		panic(err)
	}

	lines, err := ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	faces = make([]face.Face, 0, len(lines))
	for _, l := range lines {
		f, err := Base64ToFaceFeature(l)
		if err != nil {
			log.Printf("can't convert to face: %v", err)

			continue
		}

		faces = append(faces, f)
	}

	matched := RecognizeFace(dataDir, target, faces)

	fmt.Printf("matched: %d\n", matched)
}

func ExtractFaceFeature(srcDir string) ([]face.Face, error) {
	rec, err := face.NewRecognizer(srcDir)
	if err != nil {
		fmt.Println("Cannot initialize recognizer")

		return nil, fmt.Errorf("cannot initialize recognizer %w", err)
	}
	defer rec.Close()

	const srcs = 2

	result := make([]face.Face, 0, srcs)

	for i := 0; i < srcs; i++ {
		origin := filepath.Join(dataDir, fmt.Sprintf("src-%d.jpg", i))

		faces, err := rec.RecognizeFile(origin)
		if err != nil {
			log.Printf("Can't recognize: %v", err)

			continue
		}
		if len(faces) == 0 {
			log.Printf("No faces found")

			continue
		}

		result = append(result, faces[0])
	}

	return result, nil
}

func FaceFeatureToBase64(src face.Face) (string, error) {
	jsonData, err := json.Marshal(src)
	if err != nil {
		return "", fmt.Errorf("can't marshal: %w", err)
	}

	enc := base64.StdEncoding.EncodeToString(jsonData)

	return enc, nil
}

func Base64ToFaceFeature(src string) (face.Face, error) {
	var f face.Face

	dec, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return f, fmt.Errorf("can't decode: %w", err)
	}

	if err := json.Unmarshal(dec, &f); err != nil {
		return f, fmt.Errorf("can't unmarshal: %w", err)
	}

	return f, nil
}

func WriteFile(filePath string, srcs []string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("can't create file: %w", err)
	}

	defer f.Close()

	for _, s := range srcs {
		_, err := f.WriteString(s)
		f.WriteString("\n")
		if err != nil {
			return fmt.Errorf("can't write to file: %w", err)
		}
	}

	return nil
}

func ReadFile(filePath string) ([]string, error) {
	readFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	defer readFile.Close()

	return fileLines, nil
}

func RecognizeFace(dataDir string, targetFile string, orgs []face.Face) int {
	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		fmt.Println("Cannot initialize recognizer")
	}
	defer rec.Close()

	target := filepath.Join(dataDir, targetFile)
	image, err := rec.RecognizeSingleFile(target)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if image == nil {
		log.Fatalf("Not a single face on the image")
	}

	count := 0

	for _, org := range orgs {
		rec.SetSamples([]face.Descriptor{org.Descriptor}, []int32{int32(0)})

		match := rec.ClassifyThreshold(image.Descriptor, tolerance)

		if match >= 0 {
			count = count + 1
		}
	}

	return count
}

func Recognize() {
	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		fmt.Println("Cannot initialize recognizer")
	}
	defer rec.Close()

	target := filepath.Join(dataDir, "miu-4.jpg")
	image, err := rec.RecognizeSingleFile(target)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if image == nil {
		log.Fatalf("Not a single face on the image")
	}

	for i := 0; i < 2; i++ {
		origin := filepath.Join(dataDir, fmt.Sprintf("src-%d.jpg", i))

		faces, err := rec.RecognizeFile(origin)
		if err != nil {
			log.Fatalf("Can't recognize: %v", err)
		}
		if len(faces) == 0 {
			log.Fatalf("No faces found")
		}

		var mius []face.Descriptor
		var miu []int32

		for i, f := range faces {
			mius = append(mius, f.Descriptor)
			miu = append(miu, int32(i))
		}

		rec.SetSamples(mius, miu)

		match := rec.ClassifyThreshold(image.Descriptor, tolerance)
		if match < 0 {
			log.Println("Can't classify")

			continue
		}

		fmt.Println("image match")
	}
}
