package asciiArt

import (
	"fmt"
	"os"
	"strings"
)

func (art *ArtObjects) ReverseAscii() error {
	err := art.getFsForCompare()
	if err != nil {
		return err
	}
	data, err := os.ReadFile(art.OptionArg)
	if err != nil {
		return err
	}
	arr := strings.Split(string(data), "\n")

	mapAscii := make(map[rune]int)
	j := 1
	for i := ' '; i <= '~'; i++ {
		mapAscii[i] = j
		j += 9
	}
	text := ""
	for j := 0; j < len(arr); j += 8 {

		prevLetter := 0
		for i, el := range arr[j] {
			var valid bool
			if j < len(arr)-8 && el == ' ' {
				for c := 1; c < 8; c++ {
					if byte(el) != arr[j+c][i] {
						valid = true
						break
					}
				}
				if !valid {

					for f := ' '; f <= '~'; f++ {
						as := true
						for a := 0; a < 8; a++ {
							if arr[j+a][prevLetter:i+1] != art.Reverse.Standard[(mapAscii[f])+a] {
								as = false

							}
						}
						if as {
							prevLetter = i + 1
							text += string(f)
						}

					}
				}
			}
		}
		text += "\n"

	}
	fmt.Println(text)
	return nil
}

func (art *ArtObjects) getFsForCompare() error {

	std, err := os.ReadFile(
		"pkg/fontsAsciiArt/standard.txt",
	)
	if err != nil {
		return err
	}
	sh, err := os.ReadFile(
		"pkg/fontsAsciiArt/shadow.txt",
	)
	if err != nil {
		return err
	}
	thi, err := os.ReadFile(
		"pkg/fontsAsciiArt/thinkertoy.txt",
	)
	if err != nil {
		return err
	}

	standard := strings.ReplaceAll(string(std), "\r", "")
	shadow := strings.ReplaceAll(string(sh), "\r", "")
	thinkertoy := strings.ReplaceAll(string(thi), "\r", "")

	rev := ReverseAscii{
		Standard: strings.Split(
			standard,
			string('\n'),
		),
		Shadow: strings.Split(
			shadow,
			string('\n'),
		),
		Thinkertoy: strings.Split(
			thinkertoy,
			string('\n'),
		),
	}
	art.Reverse = rev
	return nil
}
