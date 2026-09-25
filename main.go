package main

import (
	"fmt"

	"github.com/joshw34/navitui/internal/cache/cache_sqlite"
	"github.com/joshw34/navitui/internal/client/navidrome"
	"github.com/joshw34/navitui/internal/controller"
	"github.com/joshw34/navitui/internal/player/mpv"
	"github.com/joshw34/navitui/internal/startup"
	"github.com/joshw34/navitui/internal/types"
	"github.com/joshw34/navitui/internal/ui"
)

func main() {
	logger := types.LoggerSetup()
	defer logger.Close()

	var err error
	var srv types.Server
	if srv, _, err = startup.RunStartup(serverFactory, logger); err != nil {
		logger.Both("Startup failed: %v", err)
		return
	}

	var db types.Cache
	if db, err = cache_sqlite.New(logger); err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	var play types.Player
	if play, err = mpv.New(); err != nil {
		fmt.Println(err)
		return
	}
	defer play.Close()

	var ctrl *controller.Controller
	if ctrl, err = controller.New(srv, db, play); err != nil {
		fmt.Println(err)
		return
	}

	if err = ui.StartUI(ctrl); err != nil {
		fmt.Println(err)
		return
	}
}

func serverFactory(serverType string, cred startup.Credentials, logger *types.NavituiLogger) (types.Server, error) {
	switch serverType {
	case "navidrome":
		return navidrome.New(cred.BaseUrl, cred.User, cred.Password, logger)
	}
	return nil, fmt.Errorf("invalid server type: %s", serverType)
}
