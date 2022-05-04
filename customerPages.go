package main

import (
	"net/http"
)

func booking(res http.ResponseWriter, req *http.Request) {
	myCustomer := getUserCustomer(res, req)
	tpl.ExecuteTemplate(res, "booking.gohtml", myCustomer)
}

func search(res http.ResponseWriter, req *http.Request) {
	myCustomer := getUserCustomer(res, req)
	tpl.ExecuteTemplate(res, "search.gohtml", myCustomer)

}
