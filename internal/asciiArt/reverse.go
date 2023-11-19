package asciiArt

import (
	"errors"
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
	text, err := FUCK(arr, art.Reverse.Standard, mapAscii)
	if err == nil {
		art.Result = text[:len(text)-1]
		return nil
	}
	text, err = FUCK(arr, art.Reverse.Shadow, mapAscii)
	if err == nil {
		art.Result = text[:len(text)-1]
		return nil
	}
	text, err = FUCK(arr, art.Reverse.Thinkertoy, mapAscii)
	if err == nil {
		art.Result = text[:len(text)-1]
		return nil
	}

	return err
}
func FUCK(arr, style []string, mapAscii map[rune]int) (string, error) {
	text := ""

	for j := 0; j < len(arr); j += 8 {
		if len(arr[j]) == 0 {
			if j != len(arr)-1 {
				text += "\n"
			}
			j -= 7

			continue
		}
		counter := 0
		for i := 0; i < 8; i++ {
			counter += len(arr[j+i])
		}
		if counter/8 != len(arr[j]) {
			return "", errors.New(LimitationsConvert)
		}
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
							if arr[j+a][prevLetter:i+1] != style[(mapAscii[f])+a] {
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
		if j != len(arr)-1 {
			text += "\n"
		}

	}
	return text, nil
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
