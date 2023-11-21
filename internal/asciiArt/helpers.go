package asciiArt

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode"
)

func (art *ArtObjects) GetDatas() error {
	// defer art.GetTerminalVids()

	var err error

	switch art.Args[len(art.Args)-1] {
	case Standard:
		err = art.GetFs(true, Standard)
	case Shadow:
		err = art.GetFs(true, Shadow)
	case Thinkertoy:
		err = art.GetFs(true, Thinkertoy)
	default:
		if len(art.Args) < 2 {
			err = art.GetFs(false, Standard)
		} else {
			return errors.New(ExpectedStyle)
		}
	}
	if err != nil {
		return err
	}

	art.Text = strings.ReplaceAll(art.Args[len(art.Args)-1], "\\n", string('\n'))
	art.Args = art.Args[:len(art.Args)-1]

	if len(art.Args) > 1 {
		return errors.New(ExpectedArgs)
	}
	err = IsEngByLoop(art.Text)
	if err != nil {
		return errors.New(IncorectLang)
	}
	return nil
}

func (art *ArtObjects) GetOption() (error, bool) {
	switch {
	case strings.HasPrefix(art.Args[0], Align):
		err := art.GetTerminalVids()
		if err != nil {
			return err, false
		}
		if len(art.Args) > 3 || len(art.Args) < 2 {
			return errors.New(ExpectedArgs), false
		}
		art.OptionArg = strings.TrimPrefix(art.Args[0], Align)
		art.Option = Align
		art.Args = art.Args[1:]
	case strings.HasPrefix(art.Args[0], Сolour):
		if len(art.Args) > 4 || len(art.Args) < 2 {
			return errors.New(ExpectedArgs), false
		}
		art.OptionArg = strings.TrimPrefix(art.Args[0], Сolour)
		art.Option = Сolour
		art.ColorFill = []rune(art.Args[1])
		switch len(art.Args) {
		case 4:
			art.Args = art.Args[2:]
		case 3:
			switch art.Args[len(art.Args)-1] {
			case Thinkertoy, Standard, Shadow:
				art.Args = art.Args[1:]
			default:
				art.Args = art.Args[2:]

			}
		default:
			art.Args = art.Args[1:]

		}

	case strings.HasPrefix(art.Args[0], Reverse):
		if len(art.Args) != 1 {
			return errors.New(ExpectedArgs), false
		}
		art.OptionArg = strings.TrimPrefix(art.Args[0], Reverse)
		art.Option = Reverse
		return nil, false
	case strings.HasPrefix(art.Args[0], Output):
		if len(art.Args) > 3 || len(art.Args) < 2 {
			return errors.New(ExpectedArgs), false
		}
		art.OptionArg = strings.TrimPrefix(art.Args[0], Output)
		art.Option = Output
		art.Args = art.Args[1:]
	default:
		if len(art.Args) > 2 {
			return errors.New(ExpectedOptions), false
		}
	}
	return nil, true
}

func (art *ArtObjects) GetFs(cut bool, fs string) error {
	fontsPath := "pkg/fontsAsciiArt/" + fs + ".txt"
	temp, err := os.ReadFile(fontsPath)
	if err != nil {
		return err
	}

	fontsAsciiArt := strings.ReplaceAll(string(temp), "\r", "")
	art.Fs = strings.Split(fontsAsciiArt, string('\n'))
	if cut {
		art.Args = art.Args[:len(art.Args)-1]
	}

	return nil
}

func IsEngByLoop(str string) error {
	for i := 0; i < len(str); i++ {
		if str[i] > unicode.MaxASCII {
			return errors.New(IncorectLang)
		}
	}
	return nil
}

func (art *ArtObjects) GetTerminalVids() error {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	width, err := strconv.Atoi(strings.Fields(string(output))[1])
	if err != nil {
		return err
	}
	art.WidthTerm = width
	return nil
}

func CreateMapAscii() map[rune]int {
	mapAscii := make(map[rune]int)
	j := 0
	for i := ' '; i <= '~'; i++ {
		mapAscii[i] = j
		j += 9
	}
	return mapAscii
}
