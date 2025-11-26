package main

import (
	"context"
	"os"
	"syscall"

	"github.com/Nadim147c/fang"
	"github.com/Nadim147c/field/cmd"
)

var Version = ""

func main() {
	err := fang.Execute(
		context.Background(),
		cmd.Command,
		fang.WithFlagTypes(),
		fang.WithNotifySignal(os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM),
		fang.WithVersion(Version),
		fang.WithoutCompletions(),
		fang.WithoutManpage(),
		fang.WithShorthandPadding(),
	)
	if err != nil {
		os.Exit(1)
	}
}
