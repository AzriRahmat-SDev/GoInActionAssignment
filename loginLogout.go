package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
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

// func main() {

// 	//customer pages

// 	http.Handle("/favicon.ico", http.NotFoundHandler())
// 	http.ListenAndServe("127.0.0.1:5221", nil)
// }

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	// bPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	// mapAdmin["admin"] = user{"admin", bPassword, "admin", "admin"}
	// mapCustomer["azri"] = userCustomer{"azri", bPassword, "Azri", "rahmat"}
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

func alreadyLoggedIn(req *http.Request) (user *User) {
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		return nil
	}
	if username, ok := mapSessions[myCookie.Value]; ok {
		user = Users[username]
	}
	return
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var myUser user
	if req.Method == http.MethodPost {
		// get form values
		username := req.FormValue("username")
		password := req.FormValue("password")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		if username != "" {
			if _, ok := mapAdmin[username]; ok {
				http.Error(res, "Username already taken", http.StatusForbidden)
				return
			}
			id := uuid.NewV4()
			myCookie := &http.Cookie{
				Name:  "myCookie",
				Value: id.String(),
			}
			http.SetCookie(res, myCookie)
			mapSessions[myCookie.Value] = username

			bPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
			if err != nil {
				http.Error(res, "Internal server error", http.StatusInternalServerError)
				return
			}
			//store the hash in the memory, map the data in "myUser" then use this to talk the server once only
			myUser = user{username, bPassword, firstname, lastname}
			mapAdmin[username] = myUser
		}
		// redirect to main index
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return

	}
	tpl.ExecuteTemplate(res, "signup.gohtml", myUser)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		//this is the user input from the form in the html file
		username := req.FormValue("username")
		password := req.FormValue("password")
		fmt.Println(username)
		// check if admin exist with username
		myAdmin, ok := mapAdmin[username]

		if !ok {
			http.Error(res, "Username and/or password do not match", http.StatusUnauthorized)
			return
		}
		// Matching of password entered
		err := bcrypt.CompareHashAndPassword(myAdmin.Password, []byte(password))
		if err != nil {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		id := uuid.NewV4()
		myCookie := &http.Cookie{
			Name:  "myCookie",
			Value: id.String(),
		}
		http.SetCookie(res, myCookie)
		mapSessions[myCookie.Value] = username
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func loginCustomer(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		// check if user exist with username
		myUserCustomer, ok := mapCustomer[username]
		if !ok {
			http.Error(res, "Username and/or password do not match", http.StatusUnauthorized)
			return
		}
		// Matching of password entered
		err := bcrypt.CompareHashAndPassword(myUserCustomer.Password, []byte(password))
		if err != nil {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		id := uuid.NewV4()
		myCookie := &http.Cookie{
			Name:  "myCookie",
			Value: id.String(),
		}
		http.SetCookie(res, myCookie)
		mapSessions[myCookie.Value] = username
		http.Redirect(res, req, "/customerPage", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "loginCustomer.gohtml", nil)
}

func customerPage(res http.ResponseWriter, req *http.Request) {
	myCustomer := getUser(res, req)
	tpl.ExecuteTemplate(res, "customerPage.gohtml", myCustomer)
}
