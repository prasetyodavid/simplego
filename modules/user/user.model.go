/*Created Simple Module User for Intitial CRUD Test*/

package user

import "github.com/prasetyodavid/simplego/config"

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
