package main

import (
	"flag"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/wlfstn/infogfx-go/igfx"
)

func main() {
	baseImg := flag.String("baseimg", "./res/WWtempTest.png", "Base image used in the background")
	baseFont := flag.String("basefont", "./res/font/RadioCanadaBig-Regular.ttf", "Font used to be drawn over the image")
	outputResult := flag.String("output", "./bin/ex2.png", "Where you want the image output")
	flag.Parse()

	templateImg, err := igfx.LoadImgLocal(*baseImg)
	if err != nil {
		log.Fatalf("Failed to load base image: %s :: %v", *baseImg, err)
	}

	djFont, err := igfx.LoadFontLocal(*baseFont, 110.0)
	subHeaderFont, err := igfx.LoadFontLocal(*baseFont, 80.0)
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}

	textI_00 := igfx.TextInput{
		Text:     "PUPUP - DEC. 31, 2024",
		X:        480,
		Y:        290,
		TextFace: subHeaderFont,
		Color:    color.White,
	}
	textI_01 := igfx.TextInput{
		Text:     "ELICIT",
		X:        480,
		Y:        650,
		TextFace: djFont,
		Color:    color.White,
	}
	textI_02 := igfx.TextInput{
		Text:     "DJ MEOWTASTIC",
		X:        480,
		Y:        900,
		TextFace: djFont,
		Color:    color.White,
	}
	textI_03 := igfx.TextInput{
		Text:     "MIDNIGHT CITY",
		X:        480,
		Y:        1160,
		TextFace: djFont,
		Color:    color.White,
	}
	textI_04 := igfx.TextInput{
		Text:     "KHALIBER",
		X:        480,
		Y:        1400,
		TextFace: djFont,
		Color:    color.White,
	}
	textI_05 := igfx.TextInput{
		Text:     "OPEN DECK",
		X:        480,
		Y:        1680,
		TextFace: djFont,
		Color:    color.White,
	}

	canoeLicense := igfx.InitializeFromImage(templateImg)
	canoeLicense.CreateCardByTemplate(textI_00, textI_01, textI_02, textI_03, textI_04, textI_05)

	file, err := os.Create(*outputResult)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	err = png.Encode(file, canoeLicense.Image)
	if err != nil {
		log.Fatalf("Failed to encode output image: %v", err)
	}

	log.Printf("Test completed successfully, output image saved as '%s'", *outputResult)
}
