package main

import (
	"fmt"
	"net"
	"os"

	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/log"
	"github.com/dump-time/antique-trade/router"
	"github.com/fvbock/endless"
)

func main() {
	// Start server gracefully
	server := endless.NewServer(global.Config.Serv.Addr, router.R)

	// daemon mode
	if global.CmdOpts.DaemonMode {
		server.BeforeBegin = func(add string) {
			// stdout pid
			pid := os.Getpid()
			log.Info(fmt.Sprintf("Deamon started: %v", pid))
		}
	}

	// Start server
	if err := server.ListenAndServe(); err != nil {
		switch err.(type) {
		case *net.OpError:
			log.Warn(err)
		default:
			log.Fatal(err)
		}
	}
}
