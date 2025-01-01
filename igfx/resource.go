package igfx

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	ErrFontLoadFailed  = errors.New("Failed to load Font")
	ErrFontParseFailed = errors.New("Failed to parse font")
	ErrFontFaceFailed  = errors.New("Failed to generate font face")
)

func LoadFontLocal(location string, size float64) (font.Face, error) {
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

var (
	ErrImageLoadFailed        = errors.New("Failed to load image file")
	ErrImageDecodeFailed      = errors.New("Failed to decode image")
	ErrImageFormatUnsupported = errors.New("Image format is not supported")
	ErrImageFetchFailed       = errors.New("Failed to fetch image from URL")
)

func LoadImgLocal(location string) (image.Image, error) {
	imgFile, err := os.Open(location)
	if err != nil {
		return nil, ErrImageLoadFailed
	}
	defer imgFile.Close()

	var imgDecode image.Image
	extension := filepath.Ext(location)
	switch extension {
	case ".png":
		imgDecode, err = png.Decode(imgFile)
	case ".jpg":
		imgDecode, err = jpeg.Decode(imgFile)
	case ".jpeg":
		imgDecode, err = jpeg.Decode(imgFile)
	default:
		err = ErrImageFormatUnsupported
	}
	if err != nil {
		return nil, ErrImageDecodeFailed
	}

	return imgDecode, err
}

func LoadImgURL(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrImageFetchFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrImageFetchFailed
	}

	contentType := resp.Header.Get("Content-Type")
	var imgDecode image.Image
	switch contentType {
	case "image/png":
		imgDecode, err = png.Decode(resp.Body)
	case "image/jpeg":
		imgDecode, err = jpeg.Decode(resp.Body)
	default:
		err = ErrImageFormatUnsupported
	}
	if err != nil {
		return nil, ErrImageDecodeFailed
	}

	return imgDecode, nil
}
