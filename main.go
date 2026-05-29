// Gofie
package main

import (
	"fmt"
	"runtime"

	"github.com/lxsh-S/gofie/src"
)

func main() {
	// lets define terminal ANSI style
	cyan := "\033[36m" // cyan
	bold := "\033[1m"  // bold text
	reset := "\033[0m" // clear all things

	// les get our data and
	// Print it accordinf to the User's OS.
	switch runtime.GOOS {
	case "darwin":

		OS := src.GetOSNameDarwin()
		username := src.GetUserNameDarwin()
		ascii := src.PrintASCII()
		fmt.Println(ascii)

		fmt.Println(" ")

		fmt.Printf("%s%s%s%s@%sDesktop:%s\n", cyan, reset, bold, username, cyan, reset)
		fmt.Printf("%s%s---------------------\n", cyan, reset)
		fmt.Printf("%s%s%sOS:%s       %s\n", cyan, reset, cyan, reset, OS)
		fmt.Println(" ")
		fmt.Println(" ")

	case "linux":

		OS := src.GetOsName()
		username := src.GetUserName()
		ascii := src.PrintASCII()
		fmt.Println(ascii)

		fmt.Println(" ")

		fmt.Printf("%s%s%s%s@%sDesktop:%s\n", cyan, reset, bold, username, cyan, reset)
		fmt.Printf("%s%s---------------------\n", cyan, reset)
		fmt.Printf("%s%s%sOS:%s       %s\n", cyan, reset, cyan, reset, OS)
		fmt.Println(" ")
		fmt.Println(" ")
	}
}
