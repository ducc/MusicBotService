package main

import (
	"../database"
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

func createTables() error {
	stmts := getCreateStatements()
	for _, stmt := range stmts {
		_, err := db.GetDB().Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}
