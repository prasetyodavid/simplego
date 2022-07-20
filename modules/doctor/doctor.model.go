package doctor

import (
	"github.com/prasetyodavid/simplego/config"
)

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
	if err := config.DB.Exec("DELETE FROM users").Error; err != nil {
		return err
	}
	return nil
}
