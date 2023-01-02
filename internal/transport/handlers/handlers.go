package handlers

import (
	"fmt"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register...", r.Method)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Log in...")
}

func MainCurrency(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add currency...")
}

func IncomeCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Income category...")
}

func CostCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Cost ategory...")
}

func IncomeValue(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Income value...")
}

func CostValue(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Cost value...")
}

func ShowMonth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show month...")
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete account...")
}

