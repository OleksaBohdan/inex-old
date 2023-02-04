package sqlQueries

import _ "embed"

var (
	//go:embed createUser.sql
	CreateUser string
)
