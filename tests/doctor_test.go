package doctor

import (
	"encoding/json"
	"net/http"
	"testing"

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

func Get_doctor(w http.ResponseWriter, r *http.Request) {
	doctors, err := GetDoctor()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(doctors)

	}
}

func GetDoctor() ([]Doctor, error) {
	var doctors []Doctor
	result := config.DB.Preload("Specializations").Find(&doctors)
	return doctors, result.Error
}

func TestGetDoctor(t *testing.T) {

	_, err := http.NewRequest("GET", "/get_doctor", nil)
	if err != nil {
		t.Fatal(err)
	}

}
