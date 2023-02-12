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

	//go:embed readUsers.sql
	ReadUsers string

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

	//go:embed readIncomeItems.sql
	ReadIncomeItems string

	//go:embed updateIncomeItem.sql
	UpdateIncomeItem string

	//go:embed deleteIncomeItem.sql
	DeleteIncomeItem string

	//go:embed createCostItem.sql
	CreateCostItem string

	//go:embed readCostItems.sql
	ReadCostItems string

	//go:embed updateCostItem.sql
	UpdateCostItem string

	//go:embed deleteCostItem.sql
	DeleteCostItem string

	//go:embed createIncome.sql
	CreateIncome string
)
