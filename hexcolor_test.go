package hexcolor

import (
    "fmt"

    "image/color"
    "testing"
)

func TestHexRgb24BitToColor(t *testing.T) {
    expectedColor := color.NRGBA{127, 127, 127, 255}
    actualColor, err := HexRgb24BitToColor("7e7f7f")
    
    if err != nil {
        t.Error(err)
    }

    if actualColor != expectedColor {
        t.Error(fmt.Sprintf("'%#v' should equal '%#v'", actualColor, expectedColor))
    }
}
