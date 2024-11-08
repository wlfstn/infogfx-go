package infogfxgo

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

var (
	ErrImageLoadFailed        = errors.New("Failed to load image file")
	ErrImageDecodeFailed      = errors.New("Failed to decode image")
	ErrImageFormatUnsupported = errors.New("Image format is not supported")
)

func ResourceLoadImage(location string) (image.Image, error) {
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
