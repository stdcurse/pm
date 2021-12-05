package output

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Info(str string) {
	fmt.Printf("[*] %s\n", str)
}

func ErrorSimple(str string) {
	color.New(color.FgRed, color.Bold).Print("[ERROR]")
	fmt.Printf(" %s\n", str)

	os.Exit(1)
}

func Check(err error, str string, die ...bool) {
	if err != nil {
		color.New(color.FgRed, color.Bold).Print("[ERROR]")
		fmt.Printf(" %s: %s\n", str, err)

		if len(die) > 0 {
			if die[0] {
				os.Exit(1)
			}
		}
	}
}
