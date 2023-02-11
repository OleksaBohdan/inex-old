package sqlQueries

import _ "embed"

var (
	//go:embed createUser.sql
	CreateUser string

	//go:embed readUser.sql
	ReadUser string

	//go:embed updateUser.sql
	UpdateUser string

	//go:embed deleteUser.sql
	DeleteUser string

	//go:embed createNote.sql
	CreateNote string

	//go:embed readNote.sql
	ReadNote string

	//go:embed updateNote.sql
	UpdateNote string

	//go:embed deleteNote.sql
	DeleteNote string

	//go:embed createIncomeItem.sql
	CreateIncomeItem string

	//go:embed readIncomeItem.sql
	ReadIncomeItem string

	//go:embed updateIncomeItem.sql
	UpdateIncomeItem string

	//go:embed deleteIncomeItem.sql
	DeleteIncomeItem string
)
