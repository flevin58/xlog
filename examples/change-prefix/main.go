package main

import (
	"fmt"

	"github.com/flevin58/xlog"
)

func main() {
	xlog.SetShortPrefixes()
	xlog.Info("Starting")
	fmt.Println("Here the program output")
	xlog.Info("Ending")
}
