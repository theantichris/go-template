package main

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/theantichris/go-template/cmd"
)

func main() {
	if err := fang.Execute(context.Background(), cmd.Execute()); err != nil {
		os.Exit(1)
	}
}
