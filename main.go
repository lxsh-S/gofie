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

	//"Not the best way to fetch data" --> "I know" :)
	switch runtime.GOOS {
	case "darwin":

		// Les get info
		OS := src.GetOSNameDarwin()
		username := src.GetUserNameDarwin()
		CPU := src.GetCPU()
		MemUsed := src.GETMemoryUsed()
		MemTotal := src.GETMemoryTotal()

		infolines := []string{
			fmt.Sprintf("%s%s%s@%s%sDesktop:%s", bold, cyan, username, reset, red, reset),
			"--------------------------------",
			fmt.Sprintf("%sOS:%s %s%s%s%s", red, reset, bold, cyan, OS, reset),

			fmt.Sprintf("%sCPU:%s %s%s%s%s", red, reset, bold, cyan, CPU, reset),
			fmt.Sprintf("%sMemory:%s %s%s%s%s", red, reset, bold, cyan, MemUsed+"/"+MemTotal, reset),
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

			// We'll safely grab the ascii and info lines instead of crashin
			if i < len(asciiLines) {
				asciiPar = asciiLines[i]
			}

			if i < len(infolines) {
				infoPar = infolines[i]
			}

			fmt.Printf("%-55s   %s\n", asciiPar, infoPar)

		}
		fmt.Println(" ")

	case "linux":

		OS := src.GetOsName()
		username := src.GetUserName()
		CPU := src.GetCPU()
		MemUsed := src.GETMemoryUsed()
		MemTotal := src.GETMemoryTotal()

		infolines := []string{
			fmt.Sprintf("%s%s@%sDesktop:%s", bold, username, red, reset),
			"------------------",
			fmt.Sprintf("%sOS:%s %s%s%s%s", red, reset, bold, cyan, OS, reset),
			fmt.Sprintf("%sCPU:%s %s%s%s%s", red, reset, bold, cyan, CPU, reset),
			fmt.Sprintf("%sMemory:%s %s%s%s%s", red, reset, bold, cyan, MemUsed+"/"+MemTotal, reset),
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

	}
}
