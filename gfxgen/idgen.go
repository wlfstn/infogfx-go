package gfxgen

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
		case image.Image:
			log.Printf("image loaded")
		case string:
			log.Printf("text field loaded %s", input)
		default:
			log.Printf("different type: %v", v)
		}
		log.Printf("potato")
	}

	return nil, nil
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
