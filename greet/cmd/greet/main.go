package main

import (
	"os"

	"github.com/ezebunandu/greet"
)

func main() {
	greet.GreetUser(os.Stdin, os.Stdout)
}
