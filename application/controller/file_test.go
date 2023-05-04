package controller

import (
	"iscool/application/model"
	"testing"
	"time"
)

func TestUploadFile(t *testing.T) {
	// Set up test data
	username := "testuser"
	folderID := "123"
	fileName := "testfile.txt"
	description := "This is a test file."

	// Create a new user and folder
	testUser := model.User{}
	folder := model.Folder{
		ID:          folderID,
		Name:        "Test Folder",
		Description: "This is a test folder.",
		CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
	}
	testUser.Folders = map[string]model.Folder{folderID: folder}

	// Add the user to the global Users map
	Users[username] = testUser

	// Call the function being tested
	result := Upload_file(username, folderID, fileName, description)

	// Check that the function returns "Success"
	if result != "Success" {
		t.Errorf("Upload_file(%q, %q, %q, %q) = %q; want %q", username, folderID, fileName, description, result, "Success")
	}

	// Check that the file was added to the folder
	if len(testUser.Folders[folderID].Files) != 1 {
		t.Errorf("Upload_file(%q, %q, %q, %q) did not add file to folder", username, folderID, fileName, description)
	}

	// Check that the file has the correct properties
	if testUser.Folders[folderID].Files[0].FileName != fileName {
		t.Errorf("Uploaded file has incorrect FileName: got %q, want %q", testUser.Folders[folderID].Files[0].FileName, fileName)
	}
	if testUser.Folders[folderID].Files[0].Extension != "txt" {
		t.Errorf("Uploaded file has incorrect Extension: got %q, want %q", testUser.Folders[folderID].Files[0].Extension, "txt")
	}
	if testUser.Folders[folderID].Files[0].Description != description {
		t.Errorf("Uploaded file has incorrect Description: got %q, want %q", testUser.Folders[folderID].Files[0].Description, description)
	}
	// Clean up test data
	delete(Users, username)
}

func TestDeleteFile(t *testing.T) {
	// Create a user with a folder and a file
	user := model.User{
		Folders: map[string]model.Folder{
			"folder1": {
				ID:   "folder1",
				Name: "My Folder",
				Files: []model.File{
					{
						FileName:  "file1.txt",
						Extension: "txt",
					},
				},
			},
		},
	}
	Users["user1"] = user

	// Call Delete_file function to delete the file
	result := Delete_file("user1", "folder1", "file1.txt")

	// Check that the file has been deleted
	if result != "Success" {
		t.Errorf("Delete_file returned unexpected result: %v", result)
	}
	if len(user.Folders["folder1"].Files) != 0 {
		t.Errorf("Delete_file did not delete the file")
	}
}
