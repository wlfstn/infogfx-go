package igfx

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	ScaleNearestNeighbor = iota
	ScaleBilinear
)

type TextInput struct {
	Text     string
	X        int
	Y        int
	TextFace font.Face
	Color    color.Color
}

type ImageInput struct {
	Image    image.Image
	Width    int
	Height   int
	XPadding int
	YPadding int
}

type GfxDesign struct {
	Image image.Image
}

func InitializeFromImage(img image.Image) *GfxDesign {
	return &GfxDesign{Image: img}
}

func (gfx *GfxDesign) DrawImage(img image.Image, scaleMethod uint8, w, h, x, y int) {
	outputResult := image.NewRGBA(gfx.Image.Bounds())
	draw.Draw(outputResult, gfx.Image.Bounds(), gfx.Image, image.Point{}, draw.Src)

	srcWidth := img.Bounds().Dx()
	srcHeight := img.Bounds().Dy()
	scaledImg := image.NewRGBA(image.Rect(0, 0, w, h))

	switch scaleMethod {
	case ScaleNearestNeighbor:
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				srcX := x * srcWidth / w
				srcY := y * srcHeight / h
				scaledImg.Set(x, y, img.At(srcX, srcY))
			}
		}

	case ScaleBilinear:
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				// Calculate source coordinates (floating-point)
				srcX := float64(x) * float64(srcWidth) / float64(w)
				srcY := float64(y) * float64(srcHeight) / float64(h)

				x0 := int(srcX)
				y0 := int(srcY)
				x1 := x0 + 1
				y1 := y0 + 1

				if x1 >= srcWidth {
					x1 = srcWidth - 1
				}
				if y1 >= srcHeight {
					y1 = srcHeight - 1
				}

				dx := srcX - float64(x0)
				dy := srcY - float64(y0)

				c00 := color.RGBAModel.Convert(img.At(x0, y0)).(color.RGBA)
				c10 := color.RGBAModel.Convert(img.At(x1, y0)).(color.RGBA)
				c01 := color.RGBAModel.Convert(img.At(x0, y1)).(color.RGBA)
				c11 := color.RGBAModel.Convert(img.At(x1, y1)).(color.RGBA)

				r0 := (1-dx)*float64(c00.R) + dx*float64(c10.R)
				g0 := (1-dx)*float64(c00.G) + dx*float64(c10.G)
				b0 := (1-dx)*float64(c00.B) + dx*float64(c10.B)
				a0 := (1-dx)*float64(c00.A) + dx*float64(c10.A)

				r1 := (1-dx)*float64(c01.R) + dx*float64(c11.R)
				g1 := (1-dx)*float64(c01.G) + dx*float64(c11.G)
				b1 := (1-dx)*float64(c01.B) + dx*float64(c11.B)
				a1 := (1-dx)*float64(c01.A) + dx*float64(c11.A)

				r := (1-dy)*r0 + dy*r1
				g := (1-dy)*g0 + dy*g1
				b := (1-dy)*b0 + dy*b1
				a := (1-dy)*a0 + dy*a1

				scaledImg.Set(x, y, color.RGBA{
					R: uint8(r + 0.5),
					G: uint8(g + 0.5),
					B: uint8(b + 0.5),
					A: uint8(a + 0.5),
				})
			}
		}

		gfx.Image = outputResult
	}

	position := image.Rect(x, y, x+w, y+h)
	draw.Draw(outputResult, position, scaledImg, image.Point{}, draw.Over)
}

func (gfx *GfxDesign) DrawText(fontFace font.Face, x, y int, text string) error {

	col := color.Black
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	d := &font.Drawer{
		Dst:  gfx.Image.(*image.RGBA),
		Src:  image.NewUniform(col),
		Face: fontFace,
		Dot:  point,
	}
	d.DrawString(text)
	return nil
}
