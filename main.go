package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/prasetyodavid/simplego/config"
)

type Doctor struct {
	ID              int64            `sql:"auto_increment" json:"-"`
	Name            string           `json:"name"`
	Age             int64            `json:"age"`
	Specializations []Specialization `json:"specializations"`
}

type Specialization struct {
	ID          int64  `sql:"auto_increment" json:"-"`
	Number      int64  `json:"number"`
	Doctor      Doctor `gorm:"foreignkey:doctor_id" json:"-"`
	DoctorID    int64  `json:"doctor_id"`
	Description string `json:"description"`
}

type User struct {
	Id    string
	Name  string
	Phone string
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

func StoreDoctor(doctors []Doctor) error {
	if err := config.DB.Create(doctors).Error; err != nil {
		return err
	}
	return nil
}

func GetDoctor() ([]Doctor, error) {
	var doctors []Doctor
	//var specialization Specialization
	result := config.DB.Preload("Specializations").Find(&doctors)
	return doctors, result.Error
}

func store_doctor(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var doctor Doctor
	var doctors []Doctor
	json.Unmarshal(reqBody, &doctor)
	doctors = append(doctors, doctor)

	if StoreDoctor(doctors) != nil { // method create doctor
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(doctors)
	}

}

func get_doctor(w http.ResponseWriter, r *http.Request) {
	doctors, err := GetDoctor()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(doctors)

	}
}

func GetSpecialization() ([]Specialization, error) {
	var specializations []Specialization
	result := config.DB.Find(&specializations)
	return specializations, result.Error
}

func get_specialization(w http.ResponseWriter, r *http.Request) {
	specializations, err := GetSpecialization()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(specializations)

	}
}

func main() {

	config.ConnectDB()
	config.DB.AutoMigrate(&User{}, &Doctor{}, &Specialization{})

	http.HandleFunc("/store_user", store_user)
	http.HandleFunc("/get_user", get_user)

	http.HandleFunc("/store_doctor", store_doctor)
	http.HandleFunc("/get_doctor", get_doctor)

	http.HandleFunc("/get_specialization", get_specialization)

	http.ListenAndServe(":8080", nil)
}
