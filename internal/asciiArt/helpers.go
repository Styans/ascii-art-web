package asciiArt

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type ArtObjects struct {
	Text      string
	Fonts     string
	Option    string
	OptionArg string
	ColorFill []rune
	Fs        []string
	Args      []string
	WidthTerm int
	Result    string
}

const (
	Сolour  = "--colour="
	Output  = "--output="
	Reverse = "--reverse"
	Align   = "--align="
)

func (art *ArtObjects) GetDatas() error {
	// defer art.GetTerminalVids()

	var err error

	switch art.Args[len(art.Args)-1] {
	case "standard":
		err = art.GetFs(true, "standard")
	case "shadow":
		err = art.GetFs(true, "shadow")
	case "thinkertoy":
		err = art.GetFs(true, "thinkertoy")
	default:
		if len(art.Args) != 2 {
			err = art.GetFs(false, "standard")
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
	return nil
}

func (art *ArtObjects) GetOption() (error, bool) {
	switch {
	case strings.HasPrefix(art.Args[0], Align):
		if len(art.Args) > 3 || len(art.Args) < 2 {
			return errors.New(ExpectedArgs), false
		}
		art.OptionArg = strings.TrimPrefix(art.Args[0], Align)
		art.Option = Align
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
	}
	return nil, true
}

func (art *ArtObjects) GetFs(cut bool, fs string) error {
	fontsPath := "pkg/fontsAsciiArt/" + fs + ".txt"
	temp, err := os.ReadFile(fontsPath)
	if err != nil {
		return err
	}
	var arr []string
	fontsAsciiArt := strings.ReplaceAll(string(temp), "\r", string(' '))
	arr = strings.Split(fontsAsciiArt, string('\n'))
	art.Fs = arr
	if cut {
		art.Args = art.Args[:len(art.Args)-1]
	}
	return nil
}

func (art *ArtObjects) GetTerminalVids() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, _ := cmd.Output()
	width, _ := strconv.Atoi(strings.Fields(string(output))[1])
	art.WidthTerm = width
}
