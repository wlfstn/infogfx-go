package infogfx

import (
	"errors"
	"image/color"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type TextInput struct {
	Text     string
	X        int
	Y        int
	TextFace font.Face
	Color    color.Color
}

var (
	ErrFontLoadFailed  = errors.New("Failed to load Font")
	ErrFontParseFailed = errors.New("Failed to parse font")
	ErrFontFaceFailed  = errors.New("Failed to generate font face")
)

func ResourceLoadFont(location string, size float64) (font.Face, error) {
	fontBytes, err := os.ReadFile(location)
	if err != nil {
		return nil, ErrFontLoadFailed
	}

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, ErrFontParseFailed
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, ErrFontFaceFailed
	}
	return face, nil
}
