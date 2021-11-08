package main

import (
	"os"

	"github.com/fredrikburman/ted-lasso-quotes-cli/tedlassoquotes"
)

func main() {
	os.Exit(tedlassoquotes.Run(os.Args[1:]))
}
