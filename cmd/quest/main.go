package main

import (
	"fmt"
	"os"

	"github.com/alcortesm/quest/internal/cobra"
)

func main() {
	err := cobra.Do(os.Args, os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
