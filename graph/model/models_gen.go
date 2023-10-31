// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Company struct {
	ID          string `json:"id"`
	Companyname string `json:"companyname"`
	Address     string `json:"address"`
	Sal         string `json:"sal"`
}

type NewCompany struct {
	Companyname string `json:"companyname"`
	Address     string `json:"address"`
	Sal         string `json:"sal"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}