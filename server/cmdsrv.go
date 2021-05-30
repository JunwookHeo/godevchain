package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JunwookHeo/godevchain/bclogger"
)

type user struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello BC")
	bclogger.INFOMSG(w)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", cmdHandler)

	return mux
}

func StartCmdSrv() {
	srv := &http.Server{Addr: ":3000", Handler: NewHttpHandler()}
	bclogger.INFOMSG("Start CmdSrv!!")
	go CloseCmdSrv(srv)

	if err := srv.ListenAndServe(); err != nil {
		bclogger.INFOMSG("Start CmdSrv Error", err)
	}

}

func CloseCmdSrv(srv *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for sig := range c {
			// sig is a ^C, handle it
			if sig == syscall.SIGTERM || sig == syscall.SIGINT {
				srv.Close()
				bclogger.INFOMSG("Closed CmdSrv!!")
				break
			}
		}
	}()
}
