package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshw34/navitui/internal/cache"
	"github.com/joshw34/navitui/internal/client"
	"github.com/joshw34/navitui/internal/controller"
	"github.com/joshw34/navitui/internal/ui"
)

func main() {
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
	ctrl := controller.New(srv, db)
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
