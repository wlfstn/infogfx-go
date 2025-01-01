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
	baseImg := flag.String("baseimg", "./res/canoe.png", "Base image used in the background")
	baseFont := flag.String("basefont", "./res/Roboto/Roboto-Bold.ttf", "Font used to be drawn over the image")
	outputResult := flag.String("output", "./test_output.png", "Where you want the image output")
	flag.Parse()

	templateImg, err := igfx.LoadImgLocal(*baseImg)
	if err != nil {
		log.Fatalf("Failed to load base image: %s :: %v", *baseImg, err)
	}

	sampleImg, err := igfx.LoadImgURL("https://iwait.club/assets/SawyerDotWolf-B6PX3jg3.jpg")
	if err != nil {
		log.Fatalf("Failed to load online image: %v", err)
	}

	sampleFont, err := igfx.LoadFontLocal(*baseFont, 24.0)
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}

	imgInput := igfx.ImageInput{
		Image:    sampleImg,
		Width:    175,
		Height:   175,
		XPadding: 12,
		YPadding: 12,
	}

	textI_01 := igfx.TextInput{
		Text:     "1337",
		X:        8,
		Y:        292,
		TextFace: sampleFont,
		Color:    color.Black,
	}
	textI_02 := igfx.TextInput{
		Text:     "2024-11-10",
		X:        8,
		Y:        340,
		TextFace: sampleFont,
		Color:    color.Black,
	}
	textI_03 := igfx.TextInput{
		Text:     "Sawyer Greythorne",
		X:        8,
		Y:        388,
		TextFace: sampleFont,
		Color:    color.Black,
	}

	canoeLicense := igfx.InitializeFromImage(templateImg)
	canoeLicense.CreateCardByTemplate(imgInput, textI_01, textI_02, textI_03)

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
