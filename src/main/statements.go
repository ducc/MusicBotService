package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	CREATE_USERS_TABLE         = "sql/ddl/1_create_users_table.sql"
	CREATE_OAUTH_TABLE         = "sql/ddl/create_oauth_table.sql"
	CREATE_PAYMENTS_TABLE      = "sql/ddl/create_payments_table.sql"
	CREATE_SUBSCRIPTIONS_TABLE = "sql/ddl/create_subscriptions_table.sql"
)

var statements = make(map[string]string)

func listFiles(dir string) []string {
	pathArray := []string{}
	err := filepath.Walk(dir, func(path string, file os.FileInfo, err error) error {
		if file.IsDir() {
			return nil
		}
		pathArray = append(pathArray, path)
		return nil
	})
	if err != nil {
		fmt.Println("Error listing files,", err)
		return nil
	}
	return pathArray
}

func loadStatement(file string) (*string, error) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	response := string(body)
	return &response, nil
}

func loadStatements(dir string) error {
	files := listFiles(dir)
	for _, file := range files {
		stmt, err := loadStatement(file)
		if err != nil {
			return err
		}
		statements[file] = *stmt
	}
	return nil
}

func getCreateStatements() []string {
	return []string{
		statements[CREATE_USERS_TABLE],
		statements[CREATE_OAUTH_TABLE],
		statements[CREATE_PAYMENTS_TABLE],
		statements[CREATE_SUBSCRIPTIONS_TABLE],
	}
}
