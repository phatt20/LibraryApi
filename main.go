package main

import (
	"github.com/phatt20/LibraryApi/config"
	"github.com/phatt20/LibraryApi/databases"
	"github.com/phatt20/LibraryApi/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(&conf.Database)
	server := server.NewEhoServer(conf, db)
	server.Start()
}
