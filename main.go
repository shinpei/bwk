package main

import (
	"fmt"
	"os"
	"strings"
)

func Usage(flags map[string]FlagAcceptor) {
	fmt.Println("Usage: ")
	fmt.Println("\t./bwk [options] [argsuments...]")
	fmt.Println("Options:")
	for opt, f := range flags {
		fmt.Printf("\t" + opt + "\t" + f.Desc() + "\n")
	}
	fmt.Println("Arguments:")
	fmt.Println("\tFiles you want to use as an input")
}

func main() {
	sFlag := NewDelimiterFlag(" ");
	vFlag := &VarFlag{}
	pFlag := &ProgFlag{}
	flags := map[string]FlagAcceptor{
		"-F": sFlag,
		"-v": vFlag,
		"-f": pFlag,
	}
	args := os.Args
	if len(args) < 2 {
		Usage(flags)
		return
	}

	for _, arg := range args[1:] {
		// let's see how's flag
		if strings.HasPrefix(arg, "-") {
			// it seems a flag, and assumed it's a single char
			acc := flags[arg[:2]]
			if acc != nil {
				acc.Accept(arg[:2])
			} else {
				// no such flag
				fmt.Println("No such flag as '" + arg[:2] + "'")
			}
		} else {
			// seems it's an argument or the last part

		}
	}
	config := Config{
		Delimiter:sFlag.PopValue(),
	}

	core := NewCore();

	core.Exec(config, "");

}
