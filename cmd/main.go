package main

import "github.com/Projects/channel-between-websocket-and-webhook-api/api"

func main() {
	r := api.New()
	r.Run(":1234")
}