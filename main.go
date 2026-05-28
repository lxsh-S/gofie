// Gofie
package main

import (
	"fmt"

	"github.com/lxsh-S/gofie/src"
)

func main() {
	//les get our data
	//
	OS := src.GetOSName()
	username := src.GetUsername()

	// lets define terminal ANSI style
	cyan := "\033[36m" // cyan
	bold := "\033[1m"  // bold text
	reset := "\033[0m" // clear all things

	fmt.Println(" ")

	fmt.Printf("%s    /\\_/\\    %s%s%s@%sDesktop:%s\n", cyan, reset, bold, username, cyan, reset) // look like[lxsh-S@Desktop]
	fmt.Printf("%s   ( o.o )   %s---------------------\n", cyan, reset)
	fmt.Printf("%s    > ^ <    %s%sOS:%s       %s\n", cyan, reset, cyan, reset, OS) // look like[os: Arch linux]
	fmt.Printf("%s   /     \\   %s\n", cyan, reset)
	fmt.Printf("%s  (_/(_/\\_)  %s\n", cyan, reset)
	fmt.Println(" ")
	fmt.Println(" ")
}
