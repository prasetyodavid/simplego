package main

import (
	"net/http"

	"github.com/prasetyodavid/simplego/config"
	doctor "github.com/prasetyodavid/simplego/modules/doctor"
	user "github.com/prasetyodavid/simplego/modules/user"
)

func main() {

	config.ConnectDB()
	config.DB.AutoMigrate(&user.User{}, &doctor.Doctor{}, &doctor.Specialization{})

	/*Routers*/
	http.HandleFunc("/store_user", user.Store_user)
	http.HandleFunc("/get_user", user.Get_user)
	http.HandleFunc("/get_specialization", doctor.Get_specialization)
	http.HandleFunc("/delete_all", doctor.Delete_all)

	// ---- Points -----
	http.HandleFunc("/store_doctor", doctor.Store_doctor)
	http.HandleFunc("/get_doctor", doctor.Get_doctor)

	http.ListenAndServe(":8080", nil)
}
