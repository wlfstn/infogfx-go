package infogfx

import (
	"image"
	"image/color"
	"image/draw"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func CardTemplate(templateImg image.Image, inputs ...interface{}) (image.Image, error) {

	//Setup image size from template image
	outputImg := image.NewRGBA(templateImg.Bounds())
	draw.Draw(outputImg, templateImg.Bounds(), templateImg, image.Point{}, draw.Src)

	for _, input := range inputs {
		switch v := input.(type) {
		case ImageInput:
			scaledImg := ScaleImageBilinear(v.Image, v.Width, v.Height)
			position := image.Rect(v.XPadding, v.YPadding, v.XPadding+v.Width, v.YPadding+v.Height)
			draw.Draw(outputImg, position, scaledImg, image.Point{}, draw.Over)
			log.Printf("image loaded and drawn at (%d, %d)", v.XPadding, v.YPadding)
		case TextInput:
			if v.Color == nil {
				v.Color = color.Black
			}
			AddTextLabel(outputImg, v.TextFace, v.X, v.Y, v.Text)
			log.Printf("text field loaded: %s", v.Text)
		default:
			log.Printf("unknown type: %v", v)
		}
	}

	return outputImg, nil
}

func AddTextLabel(img *image.RGBA, textFace font.Face, x, y int, label string) {
	col := color.Black
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: textFace,
		Dot:  point,
	}
	d.DrawString(label)
}
