package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/ashishjuyal/RestApi/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"City" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}
type CustomerHandlers struct {
	service service.CustomerService
}

/*
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello World")
}
*/

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//	customers := []Customer{
	//	{Name: "Ashish", City: "New Dehli", Zipcode: "110075"},
	//	{Name: "Rob", City: "New Dehli", Zipcode: "110075"},
	//	}
	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
