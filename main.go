package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joshw34/navitui/internal/cache/cache_sqlite"
	"github.com/joshw34/navitui/internal/client/navidrome"
	"github.com/joshw34/navitui/internal/controller"
	"github.com/joshw34/navitui/internal/player/mpv"
	"github.com/joshw34/navitui/internal/startup"
	"github.com/joshw34/navitui/internal/types"
	"github.com/joshw34/navitui/internal/ui"
)

func main() {
	f, ferr := os.OpenFile("./logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if ferr != nil {
		fmt.Println(ferr)
	}
	log.SetOutput(f)
	defer func() { _ = f.Close() }()

	srv, _, err := startup.RunStartup(serverFactory)
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := cache_sqlite.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = db.Data.Close()
	}()
	play, err := mpv.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = play.Close()
	}()
	ctrl := controller.New(srv, db, play)
	if db.SyncRequired {
		err = ctrl.ResyncLibrary()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ui.StartUI(ctrl)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func serverFactory(serverType string, cred startup.Credentials) (types.Server, error) {
	switch serverType {
	case "navidrome":
		return navidrome.New(cred.BaseUrl, cred.User, cred.Password)
	}
	return nil, fmt.Errorf("invalid server type: %s", serverType)
}
