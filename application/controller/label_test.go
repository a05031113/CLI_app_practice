package controller

import (
	"iscool/application/model"
	"testing"
	"time"
)

func TestAddLabel(t *testing.T) {
	// Create a new user for testing
	user := model.User{
		Labels:  make(map[string]model.Label),
		Folders: make(map[string]model.Folder),
	}
	Users["testUser"] = user

	// Test adding a new label
	result := Add_label("testUser", "label1", "#ff0000")
	if result != "Success" {
		t.Errorf("Expected 'Success', but got '%s'", result)
	}
	// Test adding a label with an existing name
	result = Add_label("testUser", "label1", "#00ff00")
	if result != "the label name existing" {
		t.Errorf("Expected 'the label name existing', but got '%s'", result)
	}
	// Check that the label was added correctly
	if len(user.Labels) != 1 {
		t.Errorf("Expected 1 label, but got %d", len(user.Labels))
	}
	if user.Labels["label1"].Color != "#ff0000" {
		t.Errorf("Expected label color '#ff0000', but got '%s'", user.Labels["label1"].Color)
	}

	// Clean up test data
	delete(Users, "testUser")
}

func TestGetLabels(t *testing.T) {
	// Create a user with labels for testing
	user := model.User{
		Labels: map[string]model.Label{
			"label1": {
				LabelName:   "label1",
				Color:       "red",
				CreatedTime: time.Now().Format("2006-01-02 15:04:05"),
			},
			"label2": {
				LabelName:   "label2",
				Color:       "blue",
				CreatedTime: time.Now().Format("2006-01-02 15:04:05"),
			},
		},
		Folders: make(map[string]model.Folder),
	}
	username := "testUser"
	Users[username] = user

	// Test case 1 - valid user with labels
	expectedOutput := "label1|red|" + user.Labels["label1"].CreatedTime + "|" + username + "\nlabel2|blue|" + user.Labels["label2"].CreatedTime + "|" + username + "\n"
	output := Get_labels(username)
	if output != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, output)
	}

	// Test case 2 - user does not exist
	expectedOutput = "Error - unknown user"
	output = Get_labels("nonExistentUser")
	if output != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, output)
	}

	// Test case 3 - user has no labels
	expectedOutput = "Warning - empty labels"
	user.Labels = make(map[string]model.Label)
	Users[username] = user
	output = Get_labels(username)
	if output != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, output)
	}
	// Clean up test data
	delete(Users, username)
}

func TestDelete_labels(t *testing.T) {
	// Initialize test data
	labelName := "test_label"
	username := "test_user"
	user := model.User{
		Labels: map[string]model.Label{
			labelName: {
				LabelName:   labelName,
				Color:       "blue",
				CreatedTime: "2022-05-04 11:11:11",
			},
		},
	}
	Users[username] = user

	// Test case 1: Delete an existing label
	expectedOutput := "Success"
	actualOutput := Delete_labels(username, labelName)
	if actualOutput != expectedOutput {
		t.Errorf("Test case 1 failed - expected %s but got %s", expectedOutput, actualOutput)
	}

	// Test case 2: Try to delete a non-existing label
	expectedOutput = "Error - the label name not exist"
	actualOutput = Delete_labels(username, labelName)
	if actualOutput != expectedOutput {
		t.Errorf("Test case 2 failed - expected %s but got %s", expectedOutput, actualOutput)
	}

	// Test case 3: Try to delete a label from a non-existing user
	expectedOutput = "Error - unknown user"
	actualOutput = Delete_labels("non_existing_user", labelName)
	if actualOutput != expectedOutput {
		t.Errorf("Test case 3 failed - expected %s but got %s", expectedOutput, actualOutput)
	}

	// Clean up test data
	delete(Users, username)
}

func TestAdd_folder_label(t *testing.T) {
	// create a new user and folder to test with
	user := model.User{
		Labels: map[string]model.Label{},
		Folders: map[string]model.Folder{
			"folder1": {
				ID:          "folder1",
				Name:        "Folder 1",
				Description: "A test folder",
				CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
				Label:       []string{"label1"},
				Files:       []model.File{},
			},
		},
	}

	// add the user to the global Users map
	Users["testUser"] = user

	// test adding a label to the folder
	result := Add_folder_label("testUser", "folder1", "label2")
	if result != "Success" {
		t.Errorf("Add_folder_label failed, expected 'Success', got '%s'", result)
	}

	// test adding the same label again (should fail)
	result = Add_folder_label("testUser", "folder1", "label1")
	if result != "Error - label name already exist" {
		t.Errorf("Add_folder_label failed, expected 'Error - label name already exist', got '%s'", result)
	}

	// test adding a label to a non-existent folder (should fail)
	result = Add_folder_label("testUser", "folder2", "label1")
	if result != "Error - folderID not found" {
		t.Errorf("Add_folder_label failed, expected 'Error - folder ID not exist', got '%s'", result)
	}

	// test adding a label to a non-existent user (should fail)
	result = Add_folder_label("testUser2", "folder1", "label1")
	if result != "Error - unknown user" {
		t.Errorf("Add_folder_label failed, expected 'Error - user not exist', got '%s'", result)
	}

	// Clean up test data
	delete(Users, "testUser")
}

func TestDeleteFolderLabel(t *testing.T) {
	// Create a new user and folder
	user := model.User{
		Labels:  make(map[string]model.Label),
		Folders: make(map[string]model.Folder),
	}
	folder := model.Folder{
		ID:          "123",
		Name:        "TestFolder",
		Description: "Test Description",
		CreateAt:    "2022-05-04",
		Label:       []string{"Label1", "Label2", "Label3"},
		Files:       []model.File{},
	}
	user.Folders[folder.ID] = folder
	Users["testUser"] = user

	// Delete an existing label from the folder
	errMsg := Delete_folder_label("testUser", "123", "Label2")
	if errMsg != "Success" {
		t.Errorf("Expected success, got error message: %s", errMsg)
	}

	// Check that the label was actually deleted
	folder, _ = CheckFolder(user, "123")
	if len(folder.Label) != 2 {
		t.Errorf("Expected 2 labels, got %d", len(folder.Label))
	}
	if folder.Label[0] != "Label1" {
		t.Errorf("Expected first label to be Label1, got %s", folder.Label[0])
	}
	if folder.Label[1] != "Label3" {
		t.Errorf("Expected second label to be Label3, got %s", folder.Label[1])
	}

	// Try to delete a non-existing label
	errMsg = Delete_folder_label("testUser", "123", "Label4")
	if errMsg != "label not found" {
		t.Errorf("Expected 'label not found', got %s", errMsg)
	}

	// Clean up test data
	delete(Users, "testUser")
}
