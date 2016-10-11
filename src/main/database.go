package main

import (
	"../database"
	"fmt"
)

func createDatabase() *database.Database {
	builder := database.NewURLBuilder()
	dbconf := conf.Database
	builder.Protocol = "postgres"
	builder.Ssl = dbconf.Ssl
	builder.Host = dbconf.Host
	builder.Username = dbconf.Username
	builder.Password = dbconf.Password
	builder.Database = dbconf.Database
	return builder.BuildDatabase("postgres")
}

func createTables(callback func()) {
	_, err := db.Prepare("CREATE TABLE IF NOT EXISTS test (id int, name varchar);").Exec()
	if err != nil {
		fmt.Println("Error occured whilst creating table,", err)
		return
	}
	callback()
}
