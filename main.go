package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	args := strings.Fields(" -i sample.mmd -o sample.png -t dark -b transparent")
	cmd := exec.Command("mmdc", args...)
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
