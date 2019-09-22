package main

import (
	"github.com/go_first/db"
	"github.com/go_first/server"
)

func main() {
	db.Init()
    server.Init()
}