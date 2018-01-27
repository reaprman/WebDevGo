package models

import "time"

type User struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

// capitalize to export from package
type Session struct {
	UserName     string
	LastActivity time.Time
}
