## ASCII-ART-OUTPUT


- You must follow the same instructions as in the first subject while writing the result into a file.

The file must be named by using the flag --output=\<fileName.txt>, in which --output is the flag and \<fileName.txt> is the file name which will contain the output.

- The flag must have exactly the same format as above, any other formats must return the following usage message:
```
Usage: go run ./cmd/app/asciiArtTerminal/main.go [OPTION] [STRING] [BANNER]

EX: go run ./cmd/app/asciiArtTerminal/main.go --output=<fileName.txt> something standard
```

```
$ go run ./cmd/app/asciiArtTerminal/main.go --output=banner.txt "hello" standard
$ cat -e banner.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$
$ go run ./cmd/app/asciiArtTerminal/main.go --output=banner.txt "Hello There!" shadow
$ cat -e banner.txt
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
```