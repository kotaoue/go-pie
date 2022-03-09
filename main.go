package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/kotaoue/go-eeditor"
)

var (
	darkMode    = flag.Bool("dark", false, "dark mode")
	transparent = flag.Bool("transparent", false, "transparent background")
)

func init() {
	flag.Parse()
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	editor := eeditor.NewEditor()
	b, _ := editor.Open()
	fmt.Println(string(b))

	return drawImage()
}

func drawImage() error {
	var ss []string

	ss = append(ss, "-i sample.mmd")
	ss = append(ss, "-o sample.png")
	if *darkMode {
		ss = append(ss, "-t dark")
	}
	if *transparent {
		ss = append(ss, "-b transparent")
	}

	cmd := exec.Command("mmdc", strings.Fields(strings.Join(ss, " "))...)
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
