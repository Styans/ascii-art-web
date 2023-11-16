package asciiArt

import (
	"errors"
	"fmt"
	"strings"
)

var Colors = map[string]string{
	"Red":    "\033[31m",
	"Green":  "\033[32m",
	"Yellow": "\033[33m",
	"Blue":   "\033[34m",
	"Purple": "\033[35m",
	"Cyan":   "\033[36m",
	"White":  "\033[37m",
	"Reset":  "\033[0m",
	"Orange": "\033[38;5;208m",
}

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
	// line := ""
	// var lineRes [][]string
	// var ress [][][]string
	switch art.Option {
	case Сolour:
		for _, line := range lineText {
			for i := 0; i < 9; i++ {
				for _, letters := range line {
					art.Result += mapColour[letters] + art.Fs[mapAscii[letters]+i] + Colors["Reset"]
				}
				art.Result += "\n"

			}
		}
	case Align:
		// var allLines []string
		var argLines [][]string
		var lineDatas []string
		lineTemp := ""
		var linies [][]string
		for _, line := range text {
			argLines = append(argLines, strings.Fields(line))
		}
		// fmt.Println(Textlines)
		for _, line := range argLines {
			for _, textDatas := range line {
				for _, letter := range textDatas {
					lineTemp += art.Fs[mapAscii[letter]]
					
				}
				lineDatas = append(lineDatas, lineTemp)
				lineTemp = ""
			}
			linies = append(linies, lineDatas)
			lineDatas = nil
		}
		fmt.Println(linies)
	default:
		for _, line := range lineText {
			for i := 0; i < 9; i++ {
				for _, letters := range line {
					art.Result += art.Fs[mapAscii[letters]+i]
				}
				art.Result += "\n"

			}
		}
	}

	return nil
}
