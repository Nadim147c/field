package main

import (
	"context"
	"os"
	"syscall"

	"github.com/Nadim147c/field/cmd"
	"github.com/Nadim147c/theme"
	"github.com/charmbracelet/fang"
)

var Version = ""

func main() {
	err := fang.Execute(
		context.Background(),
		cmd.Command,
		fang.WithColorSchemeFunc(theme.FangTheme),
		fang.WithNotifySignal(os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM),
		fang.WithVersion(Version),
		fang.WithoutCompletions(),
		fang.WithoutManpage(),
		fang.WithFlagTypes(),
	)
	if err != nil {
		os.Exit(1)
	}
}
