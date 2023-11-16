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

	mapAscii := make(map[rune]int)
	j := 0
	for i := ' '; i <= '~'; i++ {
		mapAscii[i] = j
		j += 9
	}

	mapColour := make(map[rune]string)
	_, check := Colors[strings.Title(art.OptionArg)]

	if !check && art.Option == Сolour {
		return errors.New(IncorectColor)
	} else {
		for _, letters := range art.ColorFill {
			mapColour[letters] = Colors[strings.Title(art.OptionArg)]
		}
	}
	switch art.Option {
	case Сolour:
		for _, line := range lineText {
			for i := 1; i < 9; i++ {
				for _, letters := range line {
					art.Result += mapColour[letters] + art.Fs[mapAscii[letters]+i] + Colors["Reset"]
				}
				art.Result += "\n"
			}
		}
	default:
		for _, line := range lineText {
			for i := 1; i < 9; i++ {
				for _, letters := range line {
					art.Result += art.Fs[mapAscii[letters]+i]
				}
				art.Result += "\n"

			}
		}
	}

	return nil
}
