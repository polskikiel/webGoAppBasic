package main

import (
	"app/shared/server"
	"app/route"
)

func main() {

	server.Run(route.Load(), nil, server.GetServer())
}
