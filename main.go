package main

import (
	"fmt"
	"os"

	"github.com/JunwookHeo/godevchain/bclogger"
	"github.com/JunwookHeo/godevchain/cli"
	"github.com/JunwookHeo/godevchain/server"
)

func main() {
	defer os.Exit(0)
	bclogger.Init(false)

	go server.StartCmdSrv()

	cmd := cli.CommandLine{}
	cmd.Run()

	fmt.Println("main end")

}

func main2() {
	bclogger.Init(false)
	server.StartCmdSrv()
}
