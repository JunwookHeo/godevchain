package main

import (
	"os"

	"github.com/JunwookHeo/godevchain/bclogger"
	"github.com/JunwookHeo/godevchain/cli"
)

func main() {
	defer os.Exit(0)
	bclogger.Init(false)

	cmd := cli.CommandLine{}
	cmd.Run()
}

func main2() {
}
