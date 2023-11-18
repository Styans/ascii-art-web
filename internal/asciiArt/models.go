package asciiArt

type ArtObjects struct {
	Text      string
	Fonts     string
	Option    string
	OptionArg string
	ColorFill []rune
	Fs        []string
	AllFs     []string
	Args      []string
	WidthTerm int
	Result    string
	Reverse   ReverseAscii
}

type ReverseAscii struct {
	Result     string
	Shadow     []string
	Standard   []string
	Thinkertoy []string
}

const (
	Сolour     = "--colour="
	Output     = "--output="
	Reverse    = "--reverse="
	Align      = "--align="
	Standard   = "standard"
	Shadow     = "shadow"
	Thinkertoy = "thinkertoy"
)

var Colors = map[string]string{
	"Red":    "\033[31m",
	"Green":  "\033[32m",
	"Yellow": "\033[33m",
	"Blue":   "\033[34m",
	"Purple": "\033[35m",
	"Cyan":   "\033[36m",
	"White":  "\033[37m",
	"Reset":  "\033[0m",
	"Orange": "\033[38;5;208m",
}
