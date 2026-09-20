package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshw34/navitui/internal/cache"
	"github.com/joshw34/navitui/internal/client"
	"github.com/joshw34/navitui/internal/controller"
	"github.com/joshw34/navitui/internal/player"
	"github.com/joshw34/navitui/internal/ui"
)

func main() {
	f, ferr := os.OpenFile("./logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if ferr != nil {
		fmt.Println(ferr)
	}
	log.SetOutput(f)
	defer func() { _ = f.Close() }()
	loadEnv()
	srv := client.New(os.Getenv("NV_URL"), os.Getenv("NV_USER"), os.Getenv("NV_PASS"))
	db, err := cache.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = db.Data.Close()
	}()
	play, err := player.New()
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

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env")
		os.Exit(1)
	}
}
