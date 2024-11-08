package infogfxgo

import (
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func ResourceLoadFont(location string, size float64) (font.Face, error) {
	fontBytes, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}
	return face, nil
}
