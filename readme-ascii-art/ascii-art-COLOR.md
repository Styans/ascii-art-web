## ASCII-ART-COLOR

The output should manipulate colors using the flag --color=\<color> \<letters to be colored>, in which --color is the flag and \<color> is the color desired by the user and \<letters to be colored> is the letter or letters that you can chose to be colored. These colors can be achieved using different notations (color code systems, like RGB, hsl, ANSI...), it is up to you to choose which one you want to use.

- You should be able to choose between coloring a single letter or a set of letters.
- If the letter is not specified, the whole `string` should be colored.
- The flag must have exactly the same format as above, any other formats must return the following usage message:
```
Usage: go run ./cmd/app/asciiArtTerminal/main.go [OPTION] [STRING]

EX: go run ./cmd/app/asciiArtTerminal/main.go --color=<color> <letters to be colored> "something"
```