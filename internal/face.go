package internal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/Kagami/go-face"
)

const (
	DataDir   = "images"
	Tolerance = 0.13
)

func ExtractFaceFeature(target string) (face.Face, error) {
	rec, err := face.NewRecognizer(DataDir)
	if err != nil {
		fmt.Println("Cannot initialize recognizer")

		return face.Face{}, fmt.Errorf("cannot initialize recognizer %w", err)
	}
	defer rec.Close()

	origin := filepath.Join(DataDir, target)

	faces, err := rec.RecognizeFile(origin)
	if err != nil {
		return face.Face{}, fmt.Errorf("can't recognize: %w", err)
	}
	if len(faces) == 0 {
		return face.Face{}, fmt.Errorf("no faces found")
	}

	return faces[0], nil
}

func FaceFeatureToBase64(src face.Face) (string, error) {
	jsonData, err := json.Marshal(src)
	if err != nil {
		return "", fmt.Errorf("can't marshal: %w", err)
	}

	enc := base64.StdEncoding.EncodeToString(jsonData)

	return enc, nil
}

func Base64ToFaceFeature(src string) (*face.Face, error) {
	var f face.Face

	dec, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, fmt.Errorf("can't decode: %w", err)
	}

	if err := json.Unmarshal(dec, &f); err != nil {
		return nil, fmt.Errorf("can't unmarshal: %w", err)
	}

	return &f, nil
}
