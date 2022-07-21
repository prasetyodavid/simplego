package main

import (
	"net/http"

	"github.com/prasetyodavid/simplego/config"
	doctor "github.com/prasetyodavid/simplego/modules/doctor"
)

func main() {

	config.ConnectDB()
	config.DB.AutoMigrate(&doctor.Doctor{}, &doctor.Specialization{})

	/* ---- Helpers Router*/
	http.HandleFunc("/get_specialization", doctor.Get_specialization)
	http.HandleFunc("/delete_all", doctor.Delete_all)

	// ---- Requirements EndPoint -----
	http.HandleFunc("/store_doctor", doctor.Store_doctor)
	http.HandleFunc("/get_doctor", doctor.Get_doctor)

	http.ListenAndServe(":8080", nil)
}
