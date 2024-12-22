package igfx

import (
	"errors"
	"image"
	"image/color"
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

func ScaleImageNN(img image.Image, newWidth, newHeight int) *image.RGBA {
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

func ScaleImageBilinear(img image.Image, newWidth, newHeight int) *image.RGBA {
	bounds := img.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()

	scaledImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			// Calculate source coordinates (floating-point)
			srcX := float64(x) * float64(srcWidth) / float64(newWidth)
			srcY := float64(y) * float64(srcHeight) / float64(newHeight)

			// Determine the surrounding pixel indices
			x0 := int(srcX)
			y0 := int(srcY)
			x1 := x0 + 1
			y1 := y0 + 1

			// Clamp the indices to avoid going out of bounds
			if x1 >= srcWidth {
				x1 = srcWidth - 1
			}
			if y1 >= srcHeight {
				y1 = srcHeight - 1
			}

			// Calculate weights for interpolation
			dx := srcX - float64(x0)
			dy := srcY - float64(y0)

			// Get pixel values for the four surrounding points
			c00 := color.RGBAModel.Convert(img.At(x0, y0)).(color.RGBA)
			c10 := color.RGBAModel.Convert(img.At(x1, y0)).(color.RGBA)
			c01 := color.RGBAModel.Convert(img.At(x0, y1)).(color.RGBA)
			c11 := color.RGBAModel.Convert(img.At(x1, y1)).(color.RGBA)

			// Interpolate in the x-direction
			r0 := (1-dx)*float64(c00.R) + dx*float64(c10.R)
			g0 := (1-dx)*float64(c00.G) + dx*float64(c10.G)
			b0 := (1-dx)*float64(c00.B) + dx*float64(c10.B)
			a0 := (1-dx)*float64(c00.A) + dx*float64(c10.A)

			r1 := (1-dx)*float64(c01.R) + dx*float64(c11.R)
			g1 := (1-dx)*float64(c01.G) + dx*float64(c11.G)
			b1 := (1-dx)*float64(c01.B) + dx*float64(c11.B)
			a1 := (1-dx)*float64(c01.A) + dx*float64(c11.A)

			// Y interpolate
			r := (1-dy)*r0 + dy*r1
			g := (1-dy)*g0 + dy*g1
			b := (1-dy)*b0 + dy*b1
			a := (1-dy)*a0 + dy*a1

			// Final Color values of scaled image
			scaledImg.Set(x, y, color.RGBA{
				R: uint8(r + 0.5),
				G: uint8(g + 0.5),
				B: uint8(b + 0.5),
				A: uint8(a + 0.5),
			})
		}
	}
	return scaledImg
}
