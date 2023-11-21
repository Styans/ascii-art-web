## ASCII-ART-REVERSE

Ascii-reverse consists on reversing the process, converting the graphic representation into a text. You will have to create a text file containing a graphic representation of a random string given as an argument.

The argument will be a flag, --reverse=\<fileName\>, in which --reverse is the flag and \<fileName\> is the file name. The program must then print this string in normal text.

- The flag must have exactly the same format as above, any other formats must return the following usage message:
```
Usage: go run ./cmd/app/asciiArtTerminal/main.go [OPTION]

EX: go run ./cmd/app/asciiArtTerminal/main.go --reverse=<fileName>
```
If there are other ascii-art optional projects implemented, the program should accept other correctly formatted [OPTION] and/or [BANNER].\
Additionally, the program must still be able to run with a single [STRING] argument.