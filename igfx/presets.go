package igfx

import (
	"image"
	"image/color"
	"image/draw"
	"log"
)

func (gfx *GfxDesign) CreateCardByTemplate(inputs ...interface{}) (image.Image, error) {

	//Setup image size from template image
	outputImg := image.NewRGBA(gfx.Image.Bounds())
	draw.Draw(outputImg, gfx.Image.Bounds(), gfx.Image, image.Point{}, draw.Src)

	for _, input := range inputs {
		switch v := input.(type) {
		case ImageInput:
			gfx.DrawImage(v.Image, ScaleBilinear, v.Width, v.Height, v.XPadding, v.YPadding)
		case TextInput:
			if v.Color == nil {
				v.Color = color.Black
			}
			gfx.DrawText(v.TextFace, v.X, v.Y, v.Text)
		default:
			log.Printf("unknown type:")
		}
	}

	return outputImg, nil
}
