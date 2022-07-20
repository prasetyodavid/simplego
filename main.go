package main

import (
    "fmt"
    "net/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
    db "database"
)

func hello(w http.ResponseWriter, req *http.Request) {
       
    // Create a new user in our database.
    db.Create(&User{
        Name: "Craig"
    })   
    
    // Find all of our users.
    var users []User
    db.Find(&users)
       
    // Output the users from the DB json encoded
    jsonEncoded, _ := json.Marshal(&users)
    fmt.Println(string(jsonEncoded))
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8080", nil)
}