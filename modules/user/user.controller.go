/*Created Simple Module User for Intitial CRUD Test*/

package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Store_user(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	var users []User
	json.Unmarshal(reqBody, &user)
	users = append(users, user)

	if StoreUser(users) != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}

}

func Get_user(w http.ResponseWriter, r *http.Request) {
	users, err := GetUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)

	}
}
