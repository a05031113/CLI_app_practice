package controller

import (
	"fmt"
	"iscool/application/model"
	"time"
)

func Add_label(username string, labelName string, color string) string {
	// Check user exist
	user, userErr := CheckUser(username)
	if userErr != nil {
		return userErr.Error()
	}
	// Check label already exist
	for name, _ := range user.Labels {
		if name == labelName {
			return "the label name existing"
		}
	}
	// Create label and add in user
	label := model.Label{
		Color:       color,
		CreatedTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	user.Labels[labelName] = label
	Users[username] = user

	return "Success"
}

func Get_labels(username string) string {
	var labels []model.Label
	// Check user exist
	user, userErr := CheckUser(username)
	if userErr != nil {
		return userErr.Error()
	}
	// Append all label into labels
	for labelName, label := range user.Labels {
		label.LabelName = labelName
		labels = append(labels, label)
	}
	// if no label
	if len(labels) == 0 {
		return "Warning - empty labels"
	}
	// Transfer to string
	var label_string []string
	for _, label := range labels {
		label_string = append(label_string, fmt.Sprintf("%s|%s|%s|%s", label.LabelName, label.Color, label.CreatedTime, username))
	}
	// Print the result
	var output string
	for _, label := range label_string {
		output += label + "\n"
	}
	return output
}

func Delete_labels(username string, labelName string) string {
	// Check user exist
	user, userErr := CheckUser(username)
	if userErr != nil {
		return userErr.Error()
	}
	// Check label exist
	_, labelExist := user.Labels[labelName]
	if !labelExist {
		return "Error - the label name not exist"
	}
	// Delete label
	delete(user.Labels, labelName)
	return "Success"
}

func Add_folder_label(username string, folderID string, labelName string) string {
	// Check user exist
	user, userErr := CheckUser(username)
	if userErr != nil {
		return userErr.Error()
	}
	// Check folder exist
	folder, folderErr := CheckFolder(user, folderID)
	if folderErr != nil {
		return folderErr.Error()
	}
	// Check label exist
	for _, label := range folder.Label {
		if label == labelName {
			return "Error - label name already exist"
		}
	}
	folder.Label = append(folder.Label, labelName)
	return "Success"
}

func Delete_folder_label(username string, folderID string, labelName string) string {
	// Check user exist
	user, userErr := CheckUser(username)
	if userErr != nil {
		return userErr.Error()
	}
	// Check folder exist
	folder, folderErr := CheckFolder(user, folderID)
	if folderErr != nil {
		return folderErr.Error()
	}
	// Delete folder label
	for i, label := range folder.Label {
		if label == labelName {
			folder.Label = append(folder.Label[:i], folder.Label[i+1:]...)
			user.Folders[folderID] = folder
			return "Success"
		}
	}
	// If didn't find out specific file
	return "label not found"
}
