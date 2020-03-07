package models

type User struct {
	ID        int
	Username  string
	Password  string
	FirstName string
	LastName  string
	Address   string
	City      string
	Zip       string
	Email     string
	Phone     string
	PhotoPath string
}
