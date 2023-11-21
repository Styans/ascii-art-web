package asciiArt

import (
	"errors"
	"strings"
)

func (art *ArtObjects) AlignAscii() error {
	switch art.OptionArg {
	case "left", "right", "center", "justify":
	default:
		return errors.New(IncorectAlign)
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
	lenSpace := 0

	var lines1 string
	var words1 string
	letters1 := ""
	for j, word := range lines {
		tmp := 0
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
						if i == 1 {
							tmp += (art.WidthTerm - lenSpace) / 2
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

				if tmp >= art.WidthTerm {
					return errors.New(LimitationsWidthTerm)
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
	art.Result = lines1
	return nil
}
