package doctor

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

func Store_doctor(w http.ResponseWriter, r *http.Request) {

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

func Get_doctor(w http.ResponseWriter, r *http.Request) {
	doctors, err := GetDoctor()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(doctors)

	}
}

func Get_specialization(w http.ResponseWriter, r *http.Request) {
	specializations, err := GetSpecialization()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(specializations)

	}
}

func Delete_all(w http.ResponseWriter, r *http.Request) {
	err := DeleteAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
