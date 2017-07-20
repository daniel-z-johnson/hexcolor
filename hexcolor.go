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

func HexRgb48BitToColor(hexString string) (color.Color, error) {
    rgb8, err := hexStringToNBytes(hexString, 6)
    if err != nil {
        return nil, err
    }

    rgb := make([]uint16, 6)
    for i := 0; i < 6; i++ {
        rgb[i] = uint16(rgb8[i])
    }

    return color.NRGBA64{(rgb[0] << 8) + rgb[1], 
        (rgb[2] << 8) + rgb[3], 
        (rgb[4] << 8) + rgb[5], 
        (1 << 16) - 1}, nil
}

type hexcolorError struct {
	message string
}

func (e *hexcolorError) Error() string {
	return e.message
}
