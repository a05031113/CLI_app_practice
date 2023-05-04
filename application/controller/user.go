package controller

import (
	"fmt"
	"iscool/application/model"
)

// Create a Users object
var Users = make(map[string]model.User)

func CheckUser(username string) (model.User, error) {
	user, userExist := Users[username]
	if !userExist {
		return model.User{}, fmt.Errorf("Error - unknown user")
	}
	return user, nil
}

func Register(username string) string {
	// Check user exist
	if _, userExist := Users[username]; userExist {
		return "Error - user already existing"
	}
	// Create model.User
	Users[username] = model.User{
		Folders: make(map[string]model.Folder),
		Labels:  make(map[string]model.Label),
	}

	return "Success"
}
