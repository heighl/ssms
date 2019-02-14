package models

type User struct {
	Id       int64
	Account  string
	Password string
	Name     string
	Type     int64
}
