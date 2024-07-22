package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/bitfield/script"
)

// counts the number of non-blank lines of Go code in a project
// it should recursively find all Go files int the tree rooted
// at the current directory and count their lines (ignoring any blank lines)
// and report the final total.

// “For example, you might run it something like this:
// loc
// You've written 719 lines of Go in this project.

func main() {
	re := regexp.MustCompile(".go$")
	non_emtpy := regexp.MustCompile(`.*\S.*$`)
	lines, err := script.FindFiles(".").MatchRegexp(re).Concat().MatchRegexp(non_emtpy).CountLines()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("You have written %d lines of Go in this project\n", lines)
}
