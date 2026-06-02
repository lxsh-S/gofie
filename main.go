// Gofie
package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/lxsh-S/gofie/src"
	"github.com/spf13/cobra" // Lets do thiss!!
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing Gofie: %v", err) // We'll throw it in err lane as ThePrimagen said :)
		os.Exit(1)
	}
}

// Base command without any subcommands (rootCmd)
var rootCmd = &cobra.Command{
	Use:     "gofie",
	Short:   "Gofie is a macos/linux only system fetch tool",
	Long:    "A lightweight system information tool written in Go, that display system fetch data instantly",
	Version: "v0.6.5",
	// When "./gofie"
	Run: func(cmd *cobra.Command, args []string) {
		FetchingEngine() // Will be more readible for me like this
	},
}

func FetchingEngine() {
	// lets define terminal ANSI style
	cyan := "\033[36m" // cyan
	bold := "\033[1m"  // bold text
	reset := "\033[0m" // clear all things
	red := "\033[31m"

	asciiBlock := src.PrintASCII()
	asciiLines := strings.Split(asciiBlock, "\n")

	//"Not the best way to fetch data" --> "I know" will see later:)
	switch runtime.GOOS {
	case "darwin":

		// Les get info
		OS := src.GetOSNameDarwin()
		username := src.GetUserNameDarwin()
		CPU := src.GetCPU()
		MemUsed := src.GETMemoryUsed()
		MemTotal := src.GETMemoryTotal()
		DiskTotal := src.GetDiskTotal()
		DiskUsed := src.GetDiskUsed()
		// GPU := src.GetGPU()

		infolines := []string{
			fmt.Sprintf("%s%s%s@%s%sDesktop:%s", bold, cyan, username, reset, red, reset),
			"--------------------------------",
			fmt.Sprintf("%sOS:%s %s%s%s%s", red, reset, bold, cyan, OS, reset),

			fmt.Sprintf("%sCPU:%s %s%s%s%s", red, reset, bold, cyan, CPU, reset),
			// fmt.Sprintf("%sGPU:%s %s%s%s%s", red, reset, bold, cyan, GPU, reset),
			fmt.Sprintf("%sMemory:%s %s%s%s%s", red, reset, bold, cyan, MemUsed+"/"+MemTotal, reset),
			fmt.Sprintf("%sDisk:%s %s%s%s%s", red, reset, bold, cyan, DiskUsed+"/"+DiskTotal, reset),
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

			fmt.Printf("%-35s   %s\n", asciiPar, infoPar)

		}
		fmt.Println(" ")

	case "linux":

		OS := src.GetOsName()
		username := src.GetUserName()
		CPU := src.GetCPU()
		MemUsed := src.GETMemoryUsed()
		MemTotal := src.GETMemoryTotal()
		DiskTotal := src.GetDiskTotal()
		DiskUsed := src.GetDiskUsed()
		// GPU := src.GetGPU()

		infolines := []string{
			fmt.Sprintf("%s%s@%sDesktop:%s", bold, username, red, reset),
			"------------------",
			fmt.Sprintf("%sOS:%s %s%s%s%s", red, reset, bold, cyan, OS, reset),
			fmt.Sprintf("%sCPU:%s %s%s%s%s", red, reset, bold, cyan, CPU, reset),
			// fmt.Sprintf("%sGPU:%s %s%s%s%s", red, reset, bold, cyan, GPU, reset),
			fmt.Sprintf("%sMemory:%s %s%s%s%s", red, reset, bold, cyan, MemUsed+"/"+MemTotal, reset),
			fmt.Sprintf("%sDisk:%s %s%s%s%s", red, reset, bold, cyan, DiskUsed+"/"+DiskTotal, reset),
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

			fmt.Printf("%-35s   %s\n", asciiPar, infoPar)
		}
		fmt.Println(" ")

	}
}
