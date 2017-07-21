package hexcolor

import (
    "fmt"

    "image/color"
    "testing"
)

func TestHexRgb24BitToColor(t *testing.T) {
    expectedColor := color.NRGBA{127, 127, 127, 255}
    actualColor, err := HexRgb24BitToColor("7f7f7f")
    
    if err != nil {
        t.Error(err)
    }

    if actualColor != expectedColor {
        t.Error(fmt.Sprintf("'%#v' should equal '%#v'", actualColor, expectedColor))
    }
}

func TestHexRgb24BitToColorWithExpectedError(t *testing.T) {
    _, err := HexRgb24BitToColor("7f7f7fff")
    
    if err == nil {
        t.Error("An error was expected but none occured")
    }
}

func TestHexRgb48BitToColor(t *testing.T) {
    expectedColor := color.NRGBA64{127, 127, 127, 65535}
    actualColor, err := HexRgb48BitToColor("007f007f007f")

    if err != nil {
        t.Error(err)
    }

    if actualColor != expectedColor {
        t.Error(fmt.Sprintf("'%#v' should equal '%#v'", actualColor, expectedColor))
    }
}
