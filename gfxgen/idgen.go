package gfxgen

import (
	"image"
	"log"
)

func CardID(inputs ...interface{}) {
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
}
