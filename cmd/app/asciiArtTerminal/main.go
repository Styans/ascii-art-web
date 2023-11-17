package main

import (
	"ascii-art/internal/asciiArt"
	"fmt"
	"os"
)

func main() {
	art := asciiArt.ArtObjects{
		Args: os.Args[1:],
	}
	if len(art.Args) > 4 || len(art.Args) < 1 {
		fmt.Println(asciiArt.ExpectedArgs)
		return
	}

	err, draw := art.GetOption()
	if err != nil {
		fmt.Println(err)
		return
	}
	if !draw {
		err = art.ReverseAscii()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	err = art.GetDatas()
	switch {
	case err != nil:
		fmt.Print(asciiArt.IncorectInput, err)
		return
	case len(art.Args) > 1:
		fmt.Print(asciiArt.IncorectInput, asciiArt.ExpectedArgs)
		return
	}
	switch art.Option {
	case asciiArt.Align:
		err = art.AlignAscii()
	default:
		err = art.DrawAscii()
	}
	if err != nil {

		fmt.Print(asciiArt.IncorectInput, err)
		return
	}
	if art.Option == asciiArt.Output {
		err = art.Output()
		if err != nil {

			fmt.Println(err)
			return
		}
	} else {
		fmt.Print(art.Result)
	}
}
