package main

import (
	"GPay/api"
)

func main() {
	r := api.SetupRouter()
	r.Run(":8080")
}
