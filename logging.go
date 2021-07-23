package main

import (
	"bufio"
	"os"

	"github.com/gookit/color"
)

func logInfo(m string) {
	color.Printf("<fg=white>[</><fg=cyan;op=bold>info</><fg=white>]</> » %s\n", m)
}

func logSuccess(m string) {
	color.Printf("<fg=white>[</><fg=green;op=bold>success</><fg=white>]</> » %s\n", m)
}

func logErr(m string) {
	color.Printf("<fg=white>[</><fg=red;op=bold>err</><fg=white>]</> » %s\n", m)
}

func logWarn(m string) {
	color.Printf("<fg=white>[</><fg=yellow;op=bold>warn</><fg=white>]</> » %s\n", m)
}

func logFatal(m string) {
	color.Printf("<fg=white>[</><fg=red;op=bold>fatal err</><fg=white>]</> » %s\n", m)
	color.Printf("<fg=red>Exiting...</>\n")
	os.Exit(0)
}

func userInput(m string) string {
	reader := bufio.NewReader(os.Stdin)
	color.Printf("<fg=white>[</><fg=cyan;op=bold>input</><fg=white>]</> %s » ", m)
	out, err := reader.ReadString('\n')
	if err != nil {
		logFatal(err.Error())
	}
	return out
}
