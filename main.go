package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City string
	State string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"
	db, err := NewDB(dir,nil)
	if err != nil {
		fmt.Println(err)
	}

	employees := []User{
		{"John","23","1234567890","Google",Address{"Bangalore","Karnataka","India","148020"}},
		{"Udai","23","1234567890","Microsoft",Address{"Bangalore","Karnataka","India","148020"}},
		{"Tushar","23","1234567890","Snive",Address{"Bangalore","Karnataka","India","148020"}},
		{"Uphar","23","1234567890","Apple",Address{"Bangalore","Karnataka","India","148020"}},
		{"Yogesh","23","1234567890","Meta",Address{"Bangalore","Karnataka","India","148020"}},
		{"Vedansh","23","1234567890","Dominate",Address{"Bangalore","Karnataka","India","148020"}},
	}

	for _, employee := range employees {
		db.write("users",employee.Name, User{
			Name:    employee.Name,
            Age:     employee.Age,
            Contact: employee.Contact,
            Company: employee.Company,
            Address: employee.Address,
		})
	}

	records, err := db.read("users")
	if err!= nil {
        fmt.Println(err)
    }
	fmt.Println(records)

	allUsers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err!= nil {
			fmt.Println(err)
		}
		allUsers = append(allUsers, employeeFound)
	}
	fmt.Println(allUsers)

	// db.delete("users", "john"); err != nil{
	// 	fmt.Println(err)
	// }
}