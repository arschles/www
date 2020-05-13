package main

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	dc := gg.NewContext(1500, 500)
	// todo, more programming

	backgroundImageFilename := "./bg.png"
	backgroundImage, err := gg.LoadImage(backgroundImageFilename)
	if err != nil {
		return errors.Wrap(err, "load background image")
	}
	dc.DrawImage(backgroundImage, 0, 0)

	backgroundImage = imaging.Fill(
		backgroundImage,
		dc.Width(),
		dc.Height(),
		imaging.Center,
		imaging.Lanczos,
	)

	margin := 20.0
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)
	dc.SetColor(color.RGBA{0, 0, 0, 204})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()

	fontPath := filepath.Join("fonts", "OpenSans-Bold.ttf")
	if err := dc.LoadFontFace(fontPath, 80); err != nil {
		return errors.Wrap(err, "load font")
	}
	dc.SetColor(color.White)
	s := "arschles.com"
	marginX := 50.0
	marginY := -10.0
	textWidth, textHeight := dc.MeasureString(s)
	x = float64(dc.Width()) - textWidth - marginX
	y = float64(dc.Height()) - textHeight - marginY
	dc.DrawString(s, x, y)

	// textColor := color.White
	// fontPath = filepath.Join("fonts", "Open_Sans", "OpenSans-Bold.ttf")
	// if err := dc.LoadFontFace(fontPath, 60); err != nil {
	// 	return errors.Wrap(err, "load Open_Sans")
	// }
	// r, g, b, _ := textColor.RGBA()
	// mutedColor := color.RGBA{
	// 	R: uint8(r),
	// 	G: uint8(g),
	// 	B: uint8(b),
	// 	A: uint8(200),
	// }
	// dc.SetColor(mutedColor)
	// marginY = 30
	// s = "arschles.com"
	// _, textHeight = dc.MeasureString(s)
	// x = 70
	// y = float64(dc.Height()) - textHeight - marginY
	// dc.DrawString(s, x, y)

	title := "All About Go Modules ... in 5 Minutes"
	textShadowColor := color.Black
	textColor := color.White
	fontPath = filepath.Join("fonts", "Open_Sans", "OpenSans-Bold.ttf")
	if err := dc.LoadFontFace(fontPath, 90); err != nil {
		return errors.Wrap(err, "load Playfair_Display")
	}
	textRightMargin := 60.0
	textTopMargin := 90.0
	x = textRightMargin
	y = textTopMargin
	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin
	dc.SetColor(textShadowColor)
	dc.DrawStringWrapped(title, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	dc.SetColor(textColor)
	dc.DrawStringWrapped(title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)

	outputFilename := "newthing.png"
	if err := dc.SavePNG(outputFilename); err != nil {
		return errors.Wrap(err, "save png")
	}
	// outputFilename := "outbg.png"
	// if _, err := os.Create(outputFilename); err != nil {
	// 	log.Fatalf("Error creating output file (%s)", err)
	// }
	// if err := dc.SavePNG(outputFilename); err != nil {
	// 	log.Fatal(err)
	// }

	// palettedImage1 := image.NewPaletted(frame1.Bounds(), palette.Plan9)
	// draw.FloydSteinberg.Draw(palettedImage1, frame1.Bounds(), frame1, image.ZP)
	// palettedImage2 := image.NewPaletted(frame2.Bounds(), palette.Plan9)
	// draw.FloydSteinberg.Draw(palettedImage2, frame2.Bounds(), frame2, image.ZP)
	// f, err := os.Create("/path/to/social-image.gif")
	// if err != nil {
	// 	return errors.Wrap(err, "create gif file")
	// }
	// gif.EncodeAll(f, &gif.GIF{
	// 	Image: []*image.Paletted{
	// 		palettedImage1,
	// 		palettedImage2,
	// 	},
	// 	Delay: []int{50, 50},
	// })
	return nil
}

// func writeString(
// 	str string,
// 	dc *gg.Context,
// 	textColor *color.Color,
// 	textShadowColor *color.Color,
// ) error {
// 	title := "All About Go Modules ... in 5 Minutes"
// 	textShadowColor := color.Black
// 	textColor := color.White
// 	fontPath = filepath.Join("fonts", "Open_Sans", "OpenSans-Bold.ttf")
// 	if err := dc.LoadFontFace(fontPath, 90); err != nil {
// 		return errors.Wrap(err, "load Playfair_Display")
// 	}
// 	textRightMargin := 60.0
// 	textTopMargin := 90.0
// 	x = textRightMargin
// 	y = textTopMargin
// 	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin
// 	dc.SetColor(textShadowColor)
// 	dc.DrawStringWrapped(title, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignLeft)
// 	dc.SetColor(textColor)
// 	dc.DrawStringWrapped(title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)

// }
