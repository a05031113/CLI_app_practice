package controller

import (
	"iscool/application/model"
	"strconv"
	"testing"
	"time"
)

func TestCheckFolder(t *testing.T) {
	// Create a test user with a folder
	testUser := model.User{
		Folders: map[string]model.Folder{
			"folder1": {
				ID:          "101",
				Name:        "Folder 1",
				Description: "Test folder",
				CreateAt:    time.Now().Format(time.RFC3339),
				Label:       []string{},
				Files:       []model.File{},
			},
		},
		Labels: map[string]model.Label{},
	}

	// Test with an existing folder
	testFolder, err := CheckFolder(testUser, "folder1")
	if err != nil {
		t.Errorf("CheckFolder failed with error: %v", err)
	}
	if testFolder.Name != "Folder 1" {
		t.Errorf("CheckFolder returned wrong folder, got %v, want %v", testFolder.Name, "Folder 1")
	}

	// Test with a non-existing folder
	_, err = CheckFolder(testUser, "folder2")
	if err == nil {
		t.Error("CheckFolder should have failed with error")
	}
}

func TestCreateFolder(t *testing.T) {
	// Create a test user
	user := model.User{
		Labels:  make(map[string]model.Label),
		Folders: make(map[string]model.Folder),
	}
	Users["testUser"] = user

	// Call the function with valid inputs
	result := Create_folder("testUser", "Test Folder", "Test Description")

	// Check if the result is a valid folder ID
	if _, err := strconv.Atoi(result); err != nil {
		t.Errorf("Create_folder returned an invalid folder ID: %s", result)
	}

	// Check if the folder was actually created
	if len(user.Folders) != 1 {
		t.Errorf("Create_folder did not create the folder")
	}

	// Check if the folder details are correct
	folder := user.Folders[result]
	if folder.Name != "Test Folder" || folder.Description != "Test Description" {
		t.Errorf("Create_folder created the folder with incorrect details")
	}
	// Clean up test data
	delete(Users, "testUser")
}

func TestDeleteFolder(t *testing.T) {
	// create a user and add a folder
	username := "test_user"
	user := model.User{Folders: make(map[string]model.Folder)}
	Users[username] = user
	folderID := Create_folder(username, "test_folder", "test description")

	// check that the folder exists
	if _, ok := user.Folders[folderID]; !ok {
		t.Fatalf("Folder %v not created for user %v", folderID, username)
	}

	// test case: user and folder exist
	result := Delete_folder(username, folderID)
	if result != "Success" {
		t.Fatalf("Unexpected result: %v", result)
	}

	// check that the folder was deleted
	if _, ok := user.Folders[folderID]; ok {
		t.Fatalf("Folder %v not deleted for user %v", folderID, username)
	}

	// test case: user does not exist
	result = Delete_folder("unknown_user", folderID)
	if result != "Error - unknown user" {
		t.Fatalf("Unexpected result: %v", result)
	}

	// test case: folder does not exist
	result = Delete_folder(username, "unknown_folder")
	if result != "Error - folder doesn't exist" {
		t.Fatalf("Unexpected result: %v", result)
	}
	// Clean up test data
	delete(Users, username)
}

func TestRenameFolder(t *testing.T) {
	// create a test user with a folder
	testUser := model.User{
		Folders: map[string]model.Folder{
			"folder1": {
				ID:          "folder1",
				Name:        "Test Folder",
				Description: "This is a test folder",
				CreateAt:    "2023-05-04 12:00:00",
			},
		},
	}

	// add test user to the global Users map
	Users["testUser"] = testUser

	// call the Rename_folder function with valid input
	result := Rename_folder("testUser", "folder1", "New Folder Name")

	// check that the function returns "Success"
	if result != "Success" {
		t.Errorf("Rename_folder returned unexpected result: %s", result)
	}

	// check that the folder was renamed
	folder := Users["testUser"].Folders["folder1"]
	if folder.Name != "New Folder Name" {
		t.Errorf("Rename_folder did not rename folder, expected 'New Folder Name' but got '%s'", folder.Name)
	}

	// call the Rename_folder function with invalid input (nonexistent user)
	result = Rename_folder("nonExistentUser", "folder1", "New Folder Name")

	// check that the function returns an error
	if result == "Success" {
		t.Error("Rename_folder returned unexpected result for nonexistent user input")
	}

	// call the Rename_folder function with invalid input (nonexistent folder)
	result = Rename_folder("testUser", "nonExistentFolder", "New Folder Name")

	// check that the function returns an error
	if result == "Success" {
		t.Error("Rename_folder returned unexpected result for nonexistent folder input")
	}
	// Clean up test data
	delete(Users, "testUser")
}
