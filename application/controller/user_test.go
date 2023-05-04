package controller

import (
	"iscool/application/model"
	"reflect"
	"testing"
)

func TestCheckUser(t *testing.T) {
	// Fake data
	testUser := model.User{
		Labels:  make(map[string]model.Label),
		Folders: make(map[string]model.Folder),
	}
	Users["testUser"] = testUser

	// Test the function with an existing user
	result, err := CheckUser("testUser")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify that the returned user is correct
	if !reflect.DeepEqual(result, testUser) {
		t.Errorf("Expected %v, but got %v", testUser, result)
	}

	// Test the function with a non-existing user
	result, err = CheckUser("nonExistingUser")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	// Clean up test data
	delete(Users, "testUser")
}

func TestRegister(t *testing.T) {
	// Test adding a new user
	username := "testRegister"
	expectedMsg := "Success"
	actualMsg := Register(username)
	if actualMsg != expectedMsg {
		t.Errorf("Expected message %q, but got %q", expectedMsg, actualMsg)
	}
	if _, ok := Users[username]; !ok {
		t.Errorf("User %q was not added to the Users map", username)
	}

	// Test adding an existing user
	expectedMsg = "Error - user already existing"
	actualMsg = Register(username)
	if actualMsg != expectedMsg {
		t.Errorf("Expected message %q, but got %q", expectedMsg, actualMsg)
	}
}
