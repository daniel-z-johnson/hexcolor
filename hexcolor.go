package hexcolor

import (
	"fmt"

	"encoding/hex"
	"image/color"
)

func hexStringToNBytes(hexString string, n int) ([]byte, error) {
    rgb, err := hex.DecodeString(hexString)
    if err != nil {
        return nil, err
    }

    if len(rgb) != n {
        return nil, &hexcolorError{fmt.Sprintf("Expected %d bytes got %d instead", n, len(rgb))}
    }

    return rgb, nil

}

func HexRgb24BitToColor(hexString string) (color.Color, error) {
	rgb, err := hexStringToNBytes(hexString, 3)
	if err != nil {
		return nil, err
	}

	return color.NRGBA{rgb[0], rgb[1], rgb[2], 255}, nil
}

type hexcolorError struct {
	message string
}

func (e *hexcolorError) Error() string {
	return e.message
}
