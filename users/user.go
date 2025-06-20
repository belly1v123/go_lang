package users

import "strings"

type User struct {
	Name    string
	Email   string
	Age     int
	Address string
	Status  bool
}

func (u User) CapatalizeUserName() string {
	return strings.ToUpper(u.Name)
} 

func NewUser(name string, email string, age int, address string, status bool) User {
	return User{
		Name:    name,
		Email:   email,
		Age:     age,
		Address: address,
		Status:  status,
	}

}
