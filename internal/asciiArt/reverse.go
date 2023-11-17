package asciiArt

import (
	"fmt"
	"os"
)

func (art *ArtObjects) ReverseAscii() error {
	_, err := os.ReadFile(art.OptionArg)
	if err != nil {
		return err
	}

	fmt.Println("reverse is not work now")
	return nil
}
