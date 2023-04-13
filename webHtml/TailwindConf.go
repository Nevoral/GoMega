package webJs

import (
	"fmt"
	"log"
	"os/exec"
)

func GenerateTailwindCss(name, inputFolder, outputFolder string, production bool) {
	var prod string = "--watch"
	if production {
		prod = "--minify"
	}
	out, err := exec.Command(name, []string{"-i", inputFolder, "-o", outputFolder, prod}...).Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(out))
}
