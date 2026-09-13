package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshw34/navitui/internal/cache"
	"github.com/joshw34/navitui/internal/client"
)

func main() {
	loadEnv()
	client := client.Init(os.Getenv("NV_URL"), os.Getenv("NV_USER"), os.Getenv("NV_PASS"))
	cache, err := cache.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cache.Data.Close()

	a, err := client.GetArtists()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = cache.UpdateArtists(a)
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
