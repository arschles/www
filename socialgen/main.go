package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pkg/errors"
)

var titleText = flag.String("title", "Testing!", "The title of the post")
var outFile = flag.String("out", "newthing.png", "Name of the output file")

func main() {
	flag.Parse()
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

	// write the site name
	dc.SetColor(color.White)
	s := "arschles.com"
	marginX := 50.0
	marginY := -10.0
	textWidth, textHeight := dc.MeasureString(s)
	x = float64(dc.Width()) - textWidth - marginX
	y = float64(dc.Height()) - textHeight - marginY
	dc.DrawString(s, x, y)

	// write the title
	if err := writeString(
		*titleText,
		// "All about Go modules ... in 5 minutes",
		fontPath,
		dc,
		color.White,
		color.Black,
	); err != nil {
		return errors.Wrap(err, "writing the title")
	}

	// save the file out
	log.Printf("Saving %s", *outFile)
	if err := dc.SavePNG(*outFile); err != nil {
		return errors.Wrap(err, "save png")
	}
	return nil
}
