package main

import (
	"fmt"
	"os"
	"strings"
)
// for making Pair of two
type FlagPhasePair struct {
	acc FlagAcceptor
	phase ArgParsingPhase
}

func Usage(flags map[string]FlagPhasePair) {
	fmt.Println("Usage: ")
	fmt.Println("\t./bwk [options] [argsuments...]")
	fmt.Println("Options:")
	for opt, f := range flags {
		fmt.Printf("\t" + opt + "\t" + f.acc.Desc() + "\n")
	}
	fmt.Println("Arguments:")
	fmt.Println("\tFiles you want to use as an input")
}
type ArgParsingPhase int
const (
DELIM_PHASE ArgParsingPhase = iota
VAR_PHASE
PROGFILE_PHASE
PROG_PHASE
)


func main() {
	sFlag := NewDelimiterFlag(" ");
	vFlag := &VarFlag{}
	pFlag := &ProgFlag{}


	flags := map[string]FlagPhasePair{
		"-F": {sFlag, DELIM_PHASE},
		"-v": {vFlag, VAR_PHASE},
		"-f": {pFlag, PROGFILE_PHASE},
	}

	args := os.Args
	if len(args) < 2 {
		Usage(flags)
		return
	}

	var phase ArgParsingPhase = PROG_PHASE
	var flagAcc FlagAcceptor;
	var prog string =""
	for _, arg := range args[1:] {
		// let's see how's flag
		if strings.HasPrefix(arg, "-") {
			// it seems a flag, and assumed it's a single char
			flagAcc = flags[arg[:2]].acc
			phase = flags[arg[:2]].phase;
			if flagAcc != nil {
				flagAcc.Accept(arg[2:])
				// phase done
				phase = PROG_PHASE
			} else {
				// no such flag
				fmt.Println("No such flag as '" + arg[:2] + "'")
				Usage(flags)
				return
			}
		} else {
			// seems it's an argument or the last part
			switch phase {
			case DELIM_PHASE:
				flagAcc.Accept(arg);
			case PROG_PHASE:
				prog = arg
			}
		}
	}

	config := NewConfig();
	config.Delimiter = sFlag.PopString();
	config.Print()
	core := NewCore();
	println("prog: '" + prog + "'")
	core.Exec(config, prog);

}
