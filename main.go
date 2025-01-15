package main

import (
	"flag"
	"fmt"

	"github.com/MyTempoesp/flick"
)

func main() {

	quietPtr := flag.Bool("q", false, "Quiet run")
	dev := flag.String("o", "/dev/ttyUSB0", "Output to specified serial device")

	flag.Parse()

	cmd := flag.Arg(0)

	quiet := *quietPtr

	if cmd == "" {

		if !quiet {
			flag.Usage()
		}

		return
	}

	fth, err := flick.NewForth(*dev)

	if err != nil {

		fmt.Println("(!) ", err)

		return
	}

	fth.Start()

	out, err := fth.Query(cmd)

	if err != nil {

		fmt.Println("(!) ", err)

		return
	}

	if !quiet {
		fmt.Println(out)
	}

	fth.Stop()
}
