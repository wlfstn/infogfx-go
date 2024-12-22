package example_test

import (
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/wlfstn/infogfx-go/igfx"
)

func TestCardTemplate(t *testing.T) {
	templateImg, err := igfx.ResourceLoadImage("./testdata/canoe.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	sampleImg, err := igfx.ResourceUrlLoadImage("https://iwait.club/assets/SawyerDotWolf-B6PX3jg3.jpg")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}

	sampleFont, err := igfx.ResourceLoadFont("./testdata/Roboto/Roboto-Bold.ttf", 24.0)
	if err != nil {
		t.Fatalf("Failed to load font: %v", err)
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

	outputImg, err := igfx.CardTemplate(templateImg, imgInput, textI_01, textI_02, textI_03)
	if err != nil {
		t.Fatalf("CardTemplate failed: %v", err)
	}

	file, err := os.Create("test_output.png")
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	err = png.Encode(file, outputImg)
	if err != nil {
		t.Fatalf("Failed to encode output image: %v", err)
	}

	if outputImg.Bounds().Empty() {
		t.Fatalf("Output image has empty bounds")
	}
	t.Log("Test completed successfully, output image saved as 'test_output.png'")
}
