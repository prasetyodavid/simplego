package doctor

import (
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

func StoreDoctor(doctors []Doctor) error {
	if err := config.DB.Create(doctors).Error; err != nil {
		return err
	}
	return nil
}

func GetDoctor() ([]Doctor, error) {
	var doctors []Doctor
	result := config.DB.Preload("Specializations").Find(&doctors)
	return doctors, result.Error
}

func GetSpecialization() ([]Specialization, error) {
	var specializations []Specialization
	result := config.DB.Find(&specializations)
	return specializations, result.Error
}

func DeleteAll() error {
	if err := config.DB.Exec("DELETE FROM doctors").Error; err != nil {
		return err
	}
	if err := config.DB.Exec("DELETE FROM specializations").Error; err != nil {
		return err
	}
	return nil
}
