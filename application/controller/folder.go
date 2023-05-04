package controller

import (
	"fmt"
	"iscool/application/model"
	"sort"
	"strconv"
	"time"
)

// Count folder ID
var IdCount int = 1001

func CheckFolder(user model.User, folderID string) (model.Folder, error) {
	folder, folderExist := user.Folders[folderID]
	if !folderExist {
		return model.Folder{}, fmt.Errorf("Error - folderID not found")
	}
	return folder, nil
}

func Create_folder(username string, folderName string, description string) string {
	// Check user exist
	user, userErr := CheckUser(username)
	if userErr != nil {
		return userErr.Error()
	}
	// Create folder
	folderID := strconv.Itoa(IdCount)
	IdCount++
	folder := model.Folder{
		Name:        folderName,
		Description: description,
		CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
		Files:       []model.File{},
	}
	user.Folders[folderID] = folder

	return folderID
}

func Delete_folder(username string, folderID string) string {
	// Check user exist
	user, userErr := CheckUser(username)
	if userErr != nil {
		return userErr.Error()
	}
	// Check folder exist
	_, folderErr := CheckFolder(user, folderID)
	if folderErr != nil {
		return "Error - folder doesn't exist"
	}
	// Delete folder
	delete(user.Folders, folderID)
	return "Success"
}

func Get_folders(username string, labelName *string, sortBy *string, orderBy *string) string {
	var folders []model.Folder
	// Check user exist
	user, userErr := CheckUser(username)
	if userErr != nil {
		return userErr.Error()
	}
	// Append all folder into slice
	for folderID, folder := range user.Folders {
		folder.ID = folderID
		folders = append(folders, folder)
	}
	// Check amount of folder
	if len(folders) == 0 {
		return "Warning - empty folders"
	}
	// Check label exist
	if labelName != nil {
		if _, labelExist := user.Labels[*labelName]; !labelExist {
			return "Error - label name not exist"
		}
	}
	// Sort folder
	if sortBy != nil && orderBy != nil {
		// Sort by name
		if *sortBy == "sort_name" {
			if *orderBy == "asc" {
				sort.Slice(folders, func(i, j int) bool {
					return folders[i].Name > folders[j].Name
				})
			} else {
				sort.Slice(folders, func(i, j int) bool {
					return folders[i].Name < folders[j].Name
				})
			}
			// Sort by time
		} else {
			if *orderBy == "asc" {
				sort.Slice(folders, func(i, j int) bool {
					return folders[i].CreateAt > folders[j].CreateAt
				})
			} else {
				sort.Slice(folders, func(i, j int) bool {
					return folders[i].CreateAt < folders[j].CreateAt
				})
			}
		}
	}
	// Transfer into string
	var folder_string []string
	for _, folder := range folders {
		if labelName != nil {
			folder_string = append(folder_string, fmt.Sprintf("%s|%s|%s|%s|%s|%s", folder.ID, *labelName, folder.Name, folder.Description, folder.CreateAt, username))
		} else {
			folder_string = append(folder_string, fmt.Sprintf("%s|%s|%s|%s|%s", folder.ID, folder.Name, folder.Description, folder.CreateAt, username))
		}
	}
	// Print result
	var output string
	for _, folder := range folder_string {
		output += folder + "\n"
	}
	return output
}

func Rename_folder(username string, folderID string, newFolderName string) string {
	// Check user exist
	user, userErr := CheckUser(username)
	if userErr != nil {
		return userErr.Error()
	}
	// Check specific folder exist
	folder, folderErr := CheckFolder(user, folderID)
	if folderErr != nil {
		return folderErr.Error()
	}
	// rename the folder
	folder.Name = newFolderName
	user.Folders[folderID] = folder

	return "Success"
}
