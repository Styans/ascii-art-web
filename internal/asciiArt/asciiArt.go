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

func (art *ArtObjects) AlignAscii() error {
	var text []string
	art.Text = strings.ReplaceAll(art.Text, "\\n", "\n")
	text = strings.Split(art.Text, "\n")
	var textarr [][]string
	for _, tmp := range text {
		textarr = append(textarr, strings.Fields(tmp))
	}

	mapAscii := make(map[rune]int)
	j := 0
	for i := ' '; i <= '~'; i++ {
		mapAscii[i] = j
		j += 9
	}

	var lines [][]string
	var words []string
	letters := ""
	for _, line := range textarr {
		for _, textTmp := range line {
			for _, letter := range textTmp {
				letters += art.Fs[mapAscii[letter]+1]
			}
			words = append(words, letters)
			letters = ""
		}
		lines = append(lines, words)
		words = nil
	}
	lenSpace := 0

	var lines1 string
	var words1 string
	letters1 := ""
	for j, word := range lines {
		for _, lenWord := range word {
			lenSpace += len(lenWord)
		}
		for i := 1; i < 9; i++ {
			for ix, word := range textarr[j] {
				if ix == 0 {
					switch art.OptionArg {
					case "center":
						for s := 0; s < (art.WidthTerm-lenSpace)/2; s++ {
							letters1 += " "
						}
					}
				}
				for _, letter := range word {
					letters1 += art.Fs[mapAscii[letter]+i]
				}
				if ix != len(textarr[j])-1 {
					switch art.OptionArg {
					case "left", "right", "center":
						letters1 += art.Fs[mapAscii[' ']+i]
					case "justify":
						for s := 0; s < (art.WidthTerm-lenSpace)/(len(textarr[j])-1); s++ {
							letters1 += " "
						}
					}
				}
			}
			words1 += letters1 + "\n"
			letters1 = ""
		}
		lenSpace = 0
		if j != len(lines)-1 {
			lines1 += words1 + "\n"
		} else {
			lines1 += words1
		}
		words1 = ""
	}
	fmt.Println(lines1)
	return nil
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
