package hexcolor

import (
	"fmt"

	"encoding/hex"
	"image/color"
)

func HexRgbToColor24bit(hexString string) (color.Color, error) {
	rgb, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	if len(rgb) != 3 {
		return nil, &hexcolorError{fmt.Sprintf("Expected three bytes got %d instead", len(rgb))}
	}

	return color.NRGBA{rgb[0], rgb[1], rgb[2], 255}, nil
}

type hexcolorError struct {
	message string
}

func (e *hexcolorError) Error() string {
	return e.message
}
