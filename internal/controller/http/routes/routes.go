package routes

import (
	"inex/main/internal/controller/http/handlers"
	"net/http"
)

func HttpRoutes() {
	// Send static html page
	fs := http.FileServer(http.Dir("../web"))
	http.Handle("/", fs)

	// sign-up user
	http.HandleFunc("/auth/sign-up", handlers.SignUp)

	// sign-in user
	http.HandleFunc("/auth/sign-in", handlers.SignIn)

	// add and change account main currency
	http.HandleFunc("/api/main-currency", handlers.MainCurrency)

	// CRUD with income category
	http.HandleFunc("/api/income-category", handlers.IncomeCategory)

	// CRUD with costs category
	http.HandleFunc("/api/cost-category", handlers.CostCategory)

	// CRUD with income value
	http.HandleFunc("/api/income-value", handlers.IncomeValue)

	// CRUD with cost value
	http.HandleFunc("/api/cost-value", handlers.CostValue)

	// get month data
	http.HandleFunc("/api/show-month", handlers.ShowMonth)

	// get month data
	http.HandleFunc("/api/delete-account", handlers.DeleteAccount)
}
