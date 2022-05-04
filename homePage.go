package main

import "net/http"

func homePage(w http.ResponseWriter, r *http.Request) {

}

func getUser(r *http.Request) (user *User) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil
	}

	if username, ok := Sessions[cookie.Value]; ok {
		user = Users[username]
	}
	return
}
