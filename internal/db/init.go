// Package db: Database handler
package db

import (
	"fmt"
	"os"
)

func getCacheDir() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("Could not fetch UserCacheDir")
		os.Exit(1)
	}
	return dir
}

func Init() {
	cacheDir := getCacheDir()
	_, err := os.Stat(cacheDir + "navitui")
}
