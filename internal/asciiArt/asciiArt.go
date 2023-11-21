package asciiArt

import (
	"errors"
	"strings"
)

func (art *ArtObjects) DrawAscii() error {
	var lineText [][]rune
	var text []string
	art.Text = strings.ReplaceAll(art.Text, "\\n", "\n")
	text = strings.Split(art.Text, "\n")
	for _, data := range text {
		lineText = append(lineText, []rune(data))
	}

	mapAscii := CreateMapAscii()

	mapColour := make(map[rune]string)
	switch art.Option {
	case Сolour:
		_, check := Colors[strings.Title(art.OptionArg)]
		if !check {
			return errors.New(IncorectColor)
		} else {
			for _, letters := range art.ColorFill {
				mapColour[letters] = Colors[strings.Title(art.OptionArg)]
			}
		}
	default:
		Colors["Reset"] = ""
	}
	art.standardAscii(lineText, mapColour, mapAscii)
	return nil
}

func (art *ArtObjects) standardAscii(lineText [][]rune, mapColour map[rune]string, mapAscii map[rune]int) {
	for _, line := range lineText {
		if len(line) == 0 {
			art.Result += "\n"
		} else {
			for i := 1; i < 9; i++ {
				for _, letters := range line {
					art.Result += mapColour[letters] + art.Fs[mapAscii[letters]+i] + Colors["Reset"]
				}
				art.Result += "\n"

			}
		}
	}
}
