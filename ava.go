package goavaname

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fogleman/gg"
)

var DEFAULT_COLOR_LIST = []color.RGBA{
	{235, 29, 54, 1},
	{255, 110, 49, 1},
	{255, 178, 0, 1},
	{139, 197, 65, 1},
	{34, 169, 224, 1},
	{60, 64, 72, 1},
	{113, 49, 221, 1},
}

func GenerateAvaname(initialName string, imgSize int, bgColor color.Color, fontColor color.Color) (image.Image, error) {

	if initialName == "" {
		return nil, errors.New("initial name can't be empty")
	}
	if len(initialName) > 2 {
		initialName = initialName[0:2]
	}

	bgImg := generateBackground(bgColor, imgSize)

	ava, err := addLabel(bgImg, labelConfiguration{
		Text:      initialName,
		Font:      "fonts/Poppins-SemiBold.ttf",
		FontSize:  float64(imgSize * 36 / 100),
		YPosition: 0,
		Color:     fontColor,
	})
	if err != nil {
		return nil, err
	}
	return ava, nil
}

func EncodeToJPEG(img image.Image, imgName string) error {
	outFile, err := os.Create(imgName)
	if err != nil {
		return err
	}
	defer outFile.Close()
	jpeg.Encode(outFile, img, &jpeg.Options{Quality: 100})
	return nil
}

func GetInitialName(value string) string {
	if value == "" {
		return ""
	}
	if strings.Contains(value, " ") {
		parts := strings.Split(value, " ")
		if len(parts) < 2 {
			return value[0:1]
		} else {
			return parts[0][0:1] + parts[1][0:1]
		}
	} else if strings.Contains(value, ".") {
		if strings.Count(value, ".") == 1 {
			if !strings.Contains(value, ".com") {
				parts := strings.Split(value, ".")
				if len(parts) < 2 {
					return value[0:1]
				} else {
					return parts[0][0:1] + parts[1][0:1]
				}
			} else {
				return value[0:1]
			}
		} else if strings.Count(value, ".") > 1 {
			parts := strings.Split(value, ".")
			if len(parts) < 2 {
				return value[0:1]
			} else {
				return parts[0][0:1] + parts[1][0:1]
			}
		}
	} else if strings.Contains(value, "_") {
		parts := strings.Split(value, "_")
		if len(parts) < 2 {
			return value[0:1]
		} else {
			return parts[0][0:1] + parts[1][0:1]
		}
	} else {
		return value[0:1]
	}
	return value[0:1]
}

func RandomColorSelector(color []color.RGBA) color.Color {
	rand.Seed(time.Now().UnixNano())
	return color[rand.Intn(len(color))]
}

func generateBackground(color color.Color, size int) *image.RGBA {
	bgImg := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(bgImg, bgImg.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)
	return bgImg
}

type labelConfiguration struct {
	Text      string
	Font      string
	FontSize  float64
	YPosition int
	Color     color.Color
}

func addLabel(bgImg *image.RGBA, label labelConfiguration) (image.Image, error) {
	bounds := bgImg.Bounds()
	imgWidth := bounds.Dx()
	imgHeight := bounds.Dy()
	dc := gg.NewContextForRGBA(bgImg)
	dc.DrawImage(bgImg, 0, 0)
	if err := dc.LoadFontFace(label.Font, label.FontSize); err != nil {
		return nil, err
	}
	x := float64(imgWidth / 2)
	y := float64((imgHeight / 2) - label.YPosition)
	maxWidth := float64(imgWidth) - 60.0
	dc.SetColor(label.Color)
	dc.DrawStringWrapped(label.Text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	return dc.Image(), nil
}
