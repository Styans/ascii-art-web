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
	
	err = art.DrawAscii()
	if err != nil {
		fmt.Print(asciiArt.IncorectInput, err)
	}
}
