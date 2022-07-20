package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/prasetyodavid/simplego/config"
)

type User struct {
	Id    string
	Name  string
	Phone string
}

type Product struct {
	ID         int64
	Name       string
	created_at int64
	updated_at int64
}

type Journal struct {
	//Journal_id       string `json:"journal_id" form:"journal_id" gorm:"primaryKey"`
	Journal_date     string `json:"journal_date" form:"journal_date"`
	Voucher_no       string `json:"voucher_no" form:"voucher_no"`
	Amount_beginning string `json:"amount_beginning" form:"amount_beginning"`
	Amount_debit     string `json:"amount_debit" form:"amount_debit"`
	Amount_credit    string `json:"amount_credit" form:"amount_credit"`
	Amount_ending    string `json:"amount_ending" form:"amount_ending"`
	Description      string `json:"description" form:"description"`
	Created_by       string `json:"stored_by" form:"stored_by"`
	Created_at       string `json:"stored_at" form:"stored_at"`
}

func model_journal() ([]Journal, error) {
	var Journal []Journal
	result := config.DB.Find(&Journal)
	return Journal, result.Error
}

func StoreUser(users []User) error {
	if err := config.DB.Create(users).Error; err != nil {
		return err
	}
	return nil
}

func GetUser() ([]User, error) {
	var User []User
	result := config.DB.Find(&User)
	return User, result.Error
}

func store_user(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	var users []User
	json.Unmarshal(reqBody, &user)
	users = append(users, user)

	if StoreUser(users) != nil { // method create user
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}

}

func get_user(w http.ResponseWriter, r *http.Request) {
	users, err := GetUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)

	}
}

func main() {

	config.ConnectDB()
	config.DB.AutoMigrate(&User{}, &Product{})

	http.HandleFunc("/store_user", store_user)
	http.HandleFunc("/get_user", get_user)

	http.ListenAndServe(":8080", nil)
}
