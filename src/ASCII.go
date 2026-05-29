package src

import (
	"fmt"
	"io"
	"os"
)

func PrintASCII() string {
	filePath := "src/data/ascii.txt"
	file, err := os.Open(filePath) // search for ascii file
	if err != nil {
		return fmt.Sprintf("error opening ASCII.txt: %v", err)
	}
	defer file.Close()

	ascii, err1 := io.ReadAll(file)
	if err != nil {
		return fmt.Sprintf("error reading ASCII.txt: %v", err1)
	}

	asciiString := string(ascii)

	return asciiString // returning string in the end.
}
