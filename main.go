package main

import (
	"os"
	"outline-go-api/pkg"
)

func main() {
	api := pkg.CreateOutlineVpn(os.Getenv("apiUrl"))
	api.GetKeys()

}
