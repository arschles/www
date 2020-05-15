package main

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/pkg/errors"
)

func writeString(
	str string,
	fontPath string,
	dc *gg.Context,
	textColor color.Color,
	textShadowColor color.Color,
) error {
	if err := dc.LoadFontFace(fontPath, 90); err != nil {
		return errors.Wrap(err, "load Playfair_Display")
	}
	textRightMargin := 60.0
	textTopMargin := 90.0
	x := textRightMargin
	y := textTopMargin
	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin
	dc.SetColor(textShadowColor)
	dc.DrawStringWrapped(str, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	dc.SetColor(textColor)
	dc.DrawStringWrapped(str, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	return nil
}
