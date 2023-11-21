package asciiArt

import (
	"errors"
	"fmt"
	"strings"
)

func (art *ArtObjects) AlignAscii() error {
	switch art.OptionArg {
	case "left", "right", "center", "justify":
	default:
		return errors.New(IncorectAlign + ExpectedAlign)
	}
	var text []string
	var textarr [][]string

	art.Text = strings.ReplaceAll(art.Text, "\\n", "\n")
	text = strings.Split(art.Text, "\n")

	for _, tmp := range text {
		textarr = append(textarr, strings.Fields(tmp))
	}

	// mapAscii := make(map[rune]int)
	mapAscii := CreateMapAscii()
	// j := 0
	// for i := ' '; i <= '~'; i++ {
	// 	mapAscii[i] = j
	// 	j += 9
	// }

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

	var lines1 string
	var words1 string
	letters1 := ""
	// fmt.Println(len(lines), lines)
	for j, word := range lines {
		lenSpace := 0
		tmp := 0
		if len(word) == 0 {
			lines1 += "\n"
			continue
		}
		for _, lenWord := range word {
			lenSpace += len(lenWord)
		}
		for i := 1; i < 9; i++ {

			for ix, word := range textarr[j] {
				if ix == 0 {
					switch art.OptionArg {
					case "center":
						for s := 0; s <= ((art.WidthTerm - lenSpace - (len(textarr[j]) * len(art.Fs[mapAscii[' ']+i]))) / 2); s++ {
							letters1 += " "
						}
					case "right":
						for s := 0; s <= (art.WidthTerm - lenSpace - (len(textarr[j]) * len(art.Fs[mapAscii[' ']+i]))); s++ {
							letters1 += " "
						}
					}
				}
				for _, letter := range word {
					letters1 += art.Fs[mapAscii[letter]+i]
					if i == 1 {
						tmp += len(string(art.Fs[mapAscii[letter]+i]))
					}
				}
				if ix != len(textarr[j])-1 {
					switch art.OptionArg {
					case "left", "right", "center":
						letters1 += art.Fs[mapAscii[' ']+i]
						if i == 1 {
							tmp += len(art.Fs[mapAscii[' ']+i])
						}

					case "justify":
						for s := 0; s < (art.WidthTerm-lenSpace)/(len(textarr[j])-1); s++ {
							letters1 += " "
						}
					}
				}
				switch art.OptionArg {
				case "center":
					fmt.Println(tmp, art.WidthTerm)
					if len(letters1) >= art.WidthTerm {
						return errors.New(LimitationsWidthTerm)
					}
				default:
					if tmp >= art.WidthTerm {
						return errors.New(LimitationsWidthTerm)
					}
				}
			}

			words1 += letters1 + "\n"
			letters1 = ""
		}
		lenSpace = 0

		lines1 += words1

		words1 = ""
	}
	art.Result = lines1
	return nil
}
