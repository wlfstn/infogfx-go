package infogfx

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
)

type ImageInput struct {
	Image    image.Image
	Width    int
	Height   int
	XPadding int
	YPadding int
}

var (
	ErrImageLoadFailed        = errors.New("Failed to load image file")
	ErrImageDecodeFailed      = errors.New("Failed to decode image")
	ErrImageFormatUnsupported = errors.New("Image format is not supported")
	ErrImageFetchFailed       = errors.New("Failed to fetch image from URL")
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

func ResourceUrlLoadImage(url string) (image.Image, error) {
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

func ScaleImage(img image.Image, newWidth, newHeight int) *image.RGBA {
	bounds := img.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()

	scaledImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := x * srcWidth / newWidth
			srcY := y * srcHeight / newHeight
			scaledImg.Set(x, y, img.At(srcX, srcY))
		}
	}
	return scaledImg
}
