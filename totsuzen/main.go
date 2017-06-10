package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/otiai10/totsuzen"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("text to decorate is needed ;)")
		os.Exit(1)
	}
	fmt.Println(totsuzen.NewToken(flag.Arg(0)).Totsuzenize().Text)
}
