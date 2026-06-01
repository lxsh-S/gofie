// Gofie
package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/lxsh-S/gofie/src"
)

func main() {
	// lets define terminal ANSI style
	cyan := "\033[36m" // cyan
	bold := "\033[1m"  // bold text
	reset := "\033[0m" // clear all things
	red := "\033[31m"

	asciiBlock := src.PrintASCII()
	asciiLines := strings.Split(asciiBlock, "\n")

	// Print it accordinf to the User's OS.
	switch runtime.GOOS {
	case "darwin":

		OS := src.GetOSNameDarwin()
		username := src.GetUserNameDarwin()
		// ascii := src.PrintASCII()
		// fmt.Println(ascii)

		infolines := []string{
			fmt.Sprintf("%s%s@%sDesktop:%s", bold, username, red, reset),
			"------------------",
			fmt.Sprintf("%sOS:%s %s%s%s%s", red, reset, bold, cyan, OS, reset),
			// fmt.Sprintf("%sCPU:%s %s%s%s%s"), //empty for now
			// fmt.Sprintf("%sMemory:%s %s%s%s%s"), //Same
		}
		fmt.Println(" ")

		// Lets check max height
		maxlines := len(asciiLines)
		if len(infolines) > maxlines {
			maxlines = len(infolines)
		}

		// Print side on side
		for i := 0; i < maxlines; i++ {
			asciiPar := ""
			infoPar := ""

			if i < len(asciiLines) {
				asciiPar = asciiLines[i]
			}

			if i < len(infolines) {
				infoPar = infolines[i]
			}

			fmt.Printf("%-55s   %s\n", asciiPar, infoPar)

		}
		fmt.Println(" ")

		// fmt.Printf("%s%s%s%s@%sDesktop:%s\n", cyan, reset, bold, username, cyan, reset)
		// fmt.Printf("%s%s---------------------\n", cyan, reset)
		// fmt.Printf("%sOS:%s%s       %s\n", cyan, reset,red, OS)
		// fmt.Println(" ")
		// fmt.Println(" ")
		//
	case "linux":

		OS := src.GetOsName()
		username := src.GetUserName()
		// ascii := src.PrintASCII()
		// fmt.Println(ascii)

		infolines := []string{
			fmt.Sprintf("%s%s@%sDesktop:%s", bold, username, red, reset),
			"------------------",
			fmt.Sprintf("%sOS:%s %s%s%s%s", red, reset, bold, cyan, OS, reset),
			// fmt.Sprintf("%sCPU:%s %s%s%s%s"), //empty for now
			// fmt.Sprintf("%sMemory:%s %s%s%s%s"), //Same
		}
		fmt.Println(" ")

		// Lets check max height
		maxlines := len(asciiLines)
		if len(infolines) > maxlines {
			maxlines = len(infolines)
		}

		// Print side on side
		for i := 0; i < maxlines; i++ {
			asciiPar := ""
			infoPar := ""

			if i < len(asciiLines) {
				asciiPar = asciiLines[i]
			}

			if i < len(infolines) {
				infoPar = infolines[i]
			}

			fmt.Printf("%-55s   %s\n", asciiPar, infoPar) // we'll print it left side and fill spaces too

		}
		fmt.Println(" ")

		// fmt.Println(" ")

		// fmt.Printf("%s%s%s%s@%sDesktop:%s\n", cyan, reset, bold, username, cyan, reset)
		// fmt.Printf("%s%s---------------------\n", cyan, reset)
		// fmt.Printf("%s%s%sOS:%s       %s\n", cyan, reset, cyan, reset, OS)
		// fmt.Println(" ")
		// fmt.Println(" ")
	}
}
