package asciiArt

import (
	"errors"
	"os"
	"strings"
)

func (art *ArtObjects) Output() error {
	if !strings.HasSuffix(art.OptionArg, ".txt") {
		return errors.New(IncorectInput + art.OptionArg)
	}
	datas, _ := os.ReadFile(art.OptionArg)

	if len(string(datas)) > 0 {
		return errors.New(IncorectInput + LimitationsFile)
	}
	f, err := os.Create(art.OptionArg)
	if err != nil {
		return err
	}

	f.WriteString(art.Result)
	return nil
}
