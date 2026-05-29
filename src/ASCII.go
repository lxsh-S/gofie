package src

import (
	_ "embed"
)

// i have noticed that there is a small delay in reading the ascii art so
// From the docs i have got to know that embed bakes the file directly

//go:embed data/ascii.txt
var asciiData string

func PrintASCII() string {
	if asciiData == "" {
		return "data/ascii.txt didn't load.." // Only when our ascii couldn't be loaded
	}
	// skip if found
	return asciiData
}
