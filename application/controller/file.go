package controller

import (
	"fmt"
	"iscool/application/model"
	"sort"
	"strings"
	"time"
)

func Upload_file(username string, folderID string, fileName string, description string) string {
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
	// Create file
	file := model.File{
		FileName:    fileName,
		Extension:   strings.Split(fileName, ".")[1],
		Description: description,
		CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
	}
	// Put file in to folder
	folder.Files = append(folder.Files, file)
	user.Folders[folderID] = folder

	return "Success"
}

func Delete_file(username string, folderID string, fileName string) string {
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
	// Select specific file and delete it
	for i, file := range folder.Files {
		if file.FileName == fileName {
			folder.Files = append(folder.Files[:i], folder.Files[i+1:]...)
			user.Folders[folderID] = folder
			return "Success"
		}
	}
	// If didn't find out specific file
	return "fileName not found"
}

func Get_files(username string, folderID string, sortBy *string, orderBy *string) string {
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
	// Check amount of files
	if len(folder.Files) == 0 {
		return "Warning - empty files"
	}
	// Sort files
	if sortBy != nil && orderBy != nil {
		// Sort by name
		if *sortBy == "sort_name" {
			if *orderBy == "asc" {
				sort.Slice(folder.Files, func(i, j int) bool {
					return folder.Files[i].FileName > folder.Files[j].FileName
				})
			} else {
				sort.Slice(folder.Files, func(i, j int) bool {
					return folder.Files[i].FileName < folder.Files[j].FileName
				})
			}
			// Sort by time
		} else if *sortBy == "sort_time" {
			if *orderBy == "asc" {
				sort.Slice(folder.Files, func(i, j int) bool {
					return folder.Files[i].CreateAt > folder.Files[j].CreateAt
				})
			} else {
				sort.Slice(folder.Files, func(i, j int) bool {
					return folder.Files[i].CreateAt < folder.Files[j].CreateAt
				})
			}
			// Sort by extension
		} else if *sortBy == "sort_extension" {
			if *orderBy == "asc" {
				sort.Slice(folder.Files, func(i, j int) bool {
					return folder.Files[i].Extension > folder.Files[j].Extension
				})
			} else {
				sort.Slice(folder.Files, func(i, j int) bool {
					return folder.Files[i].Extension < folder.Files[j].Extension
				})
			}
		}
	}
	// Transfer to string
	var filesString []string
	for _, file := range folder.Files {
		filesString = append(filesString, fmt.Sprintf("%s|%s|%s|%s|%s", file.FileName, file.Extension, file.Description, file.CreateAt, username))
	}
	// Print the result out
	var output string
	for _, file := range filesString {
		output += file + "\n"
	}
	return output
}
