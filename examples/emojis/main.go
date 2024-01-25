package main

import (
	"github.com/flevin58/xlog"
)

func main() {
	xlog.SetPrefixes("🍺 ", "😖 ", "🔥 ", "💀 ", "🙀 ")
	xlog.Info("Starting")
	xlog.Warn("Remember to initialize the variables")
	xlog.Warnf("The value is %d but should have been %d", 43, 42)
	xlog.Error("Ouch, we found an error that can't be recovered 🤯")
}
