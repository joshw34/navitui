package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/joshw34/navitui/internal/client"
	"github.com/joshw34/navitui/internal/utils"
)

/*func getURL() *url.URL {
	u, err := url.Parse(os.Getenv("NV_URL") + os.Getenv("NV_REQ"))
	utils.ErrCheck(err, "Failed to parse URL")

	query := u.Query()
	query.Set("u", os.Getenv("NV_USER"))
	query.Set("p", os.Getenv("NV_PASS"))
	query.Set("v", "1.16.1")
	query.Set("c", "navitui")
	query.Set("f", "json")
	u.RawQuery = query.Encode()

	return u
}*/

func main() {
	loadEnv()
	client := client.Init(os.Getenv("NV_URL"), os.Getenv("NV_USER"), os.Getenv("NV_PASS"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		utils.ErrorExit("Failed to load .env")
	}
}
