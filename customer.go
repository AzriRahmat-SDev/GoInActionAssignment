package main

import (
	"fmt"
	"html/template"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	customerId int
	userName   string
	Password   []byte
	firstName  string
	lastName   string
	isAdmin    bool
	bookingId  []int
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
			userName: "admin",
			Password: []byte("1234"),
			isAdmin:  true,
		}, {
			userName:  "user",
			firstName: "John",
			lastName:  "Deo",
			Password:  []byte("1234"),
			isAdmin:   false,
			bookingId: []int{1, 2, 3, 4},
		},
	}

	for _, value := range list {
		CreateNewUser(value)
	}
}

func CustomerId() int {
	max := 0
	for _, value := range Users {
		if value.customerId > max {
			max = value.customerId
		}
	}
	return max + 1
}
func CreateNewUser(u *User) error {
	u.customerId = CustomerId()

	bpassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("CreateNewUser: %w", err)
	}
	u.Password = bpassword
	Users[u.userName] = u
	return nil

}
func initilizeUsers() {
	list := []*User{
		{
			userName: "admin",
			Password: []byte("1234"),
			isAdmin:  true,
		}, {
			userName:  "user",
			firstName: "John",
			lastName:  "Doe",
			Password:  []byte("1234"),
			isAdmin:   false,
			bookingId: []int{1, 2, 3, 4, 5},
		},
	}

	for _, u := range list {
		CreateNewUser(u)
	}
}
