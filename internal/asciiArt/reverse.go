package asciiArt

import (
	"fmt"
	"os"
	"strings"
)

func (art *ArtObjects) ReverseAscii() error {
	data, err := os.ReadFile(art.OptionArg)
	if err != nil {
		return err
	}
	arrOffer := strings.Split(string(data), "\n")
	arrOffer = arrOffer
	fmt.Println("reverse is not work now")
	return nil
}
