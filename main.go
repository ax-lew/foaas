package main

import (
	"github.com/ax-lew/foaas/cmd"
	"os"
)

func main() {
	if err := cmd.Cmd().Execute(); err != nil {
		os.Exit(1)
	}
}
