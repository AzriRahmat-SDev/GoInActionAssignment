package main

import (
	"fmt"
	"html/template"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	CustomerId int
	Username   string
	Password   []byte
	Firstname  string
	Lastname   string
	isAdmin    bool
	BookingId  []int
}

var tpl *template.Template
var Users = map[string]*User{}
var Sessions = map[string]string{}

// func initCustomer() {
// 	// tpl = template.Must(template.ParseGlob("templates/*"))
// 	// bPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
// 	// mapAdmin["admin"] = user{"admin", bPassword, "admin", "admin"}
// 	// mapCustomer["azri"] = userCustomer{"azri", bPassword, "Azri", "rahmat"}
// }

func initCustomers() {
	list := []*User{
		{
			Username: "admin",
			Password: []byte("1234"),
			isAdmin:  true,
		}, {
			Username:  "user",
			Firstname: "John",
			Lastname:  "Deo",
			Password:  []byte("1234"),
			isAdmin:   false,
			BookingId: []int{1, 2, 3, 4},
		},
	}

	for _, value := range list {
		CreateNewUser(value)
	}
}

func CustomerId() int {
	max := 0
	for _, value := range Users {
		if value.CustomerId > max {
			max = value.CustomerId
		}
	}
	return max + 1
}
func CreateNewUser(u *User) error {
	u.CustomerId = CustomerId()

	bpassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("CreateNewUser: %w", err)
	}
	u.Password = bpassword
	Users[u.Username] = u
	return nil

}
func initilizeUsers() {
	list := []*User{
		{
			Username: "admin",
			Password: []byte("1234"),
			isAdmin:  true,
		}, {
			Username:  "user",
			Firstname: "John",
			Lastname:  "Doe",
			Password:  []byte("1234"),
			isAdmin:   false,
			BookingId: []int{1, 2, 3, 4, 5},
		},
	}

	for _, u := range list {
		CreateNewUser(u)
	}
}
