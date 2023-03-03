package goavaname

import (
	"image/color"
	"log"
	"testing"
)

func TestAvaname(t *testing.T) {

	name := "John Doe"
	initial := GetInitialName(name)
	imgSize := 200
	bgColor := RandomColorSelector(DEFAULT_COLOR_LIST)
	fontColor := color.White

	jpegName := "example/jd.jpeg"

	ava, err := GenerateAvaname(initial, imgSize, bgColor, fontColor)
	if err != nil {
		log.Fatal(err)
	}

	err = EncodeToJPEG(ava, jpegName)
	if err != nil {
		log.Fatal(err)
	}
}
