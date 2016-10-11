package database

import "fmt"

type URLBuilder struct {
	Protocol string
	Host     string
	Username string
	Password string
	Database string
	Ssl      bool
}

func NewURLBuilder() *URLBuilder {
	return &URLBuilder{}
}

func (builder URLBuilder) BuildURL() string {
	url := fmt.Sprintf(BASE_URL, builder.Protocol, builder.Username, builder.Password, builder.Host, builder.Database)
	if !builder.Ssl {
		url += "?sslmode=disable"
	}
	return url
}

func (builder URLBuilder) BuildDatabase(driver string) *Database {
	url := builder.BuildURL()
	return NewDatabase(driver, url)
}
