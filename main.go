package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// File structure
type File struct {
	File_name   string
	Extension   string
	Description string
	Create_at   string
}

// Folder structure
type Folder struct {
	ID          string
	Name        string
	Description string
	Create_at   string
	Files       []File
}

// User structure
type User struct {
	Folders map[string]Folder
}

// Create a Users object
var Users = make(map[string]User)

// Count folder ID
var Id_count int = 1001

func Register(username string) {
	// Check user exist
	if _, user_exist := Users[username]; user_exist {
		fmt.Println("Error - user already existing")
		return
	}
	// Create User
	Users[username] = User{
		Folders: make(map[string]Folder),
	}

	fmt.Println("Success")
	return
}

func Create_folder(username string, folder_name string, description string) {
	// Check user exist
	user, user_exist := Users[username]
	if !user_exist {
		fmt.Println("Error - unknown user")
		return
	}
	// Create folder
	folder_id := strconv.Itoa(Id_count)
	Id_count++
	folder := Folder{
		Name:        folder_name,
		Description: description,
		Create_at:   time.Now().Format("2006-01-02 15:04:05"),
		Files:       []File{},
	}
	user.Folders[folder_id] = folder

	fmt.Println(folder_id)
	return
}

func Delete_folder(username string, folder_id string) {
	// Check user exist
	user, user_exist := Users[username]
	if !user_exist {
		fmt.Println("Error - unknown user")
		return
	}
	// Check folder exist
	_, folder_exist := user.Folders[folder_id]
	if !folder_exist {
		fmt.Println("Error - folder doesn't exist")
		return
	}
	// Delete folder
	delete(user.Folders, folder_id)
	fmt.Println("Success")
	return
}

func Get_folders(username string, sort_by *string, order_by *string) {
	var folders []Folder
	// Check user exist
	user, user_exist := Users[username]
	if !user_exist {
		fmt.Println("Error - unknown user")
		return
	}
	// Append all folder into slice
	for folder_id, folder := range user.Folders {
		folder.ID = folder_id
		folders = append(folders, folder)
	}
	// Check amount of folder
	if len(folders) == 0 {
		fmt.Println("Warning - empty folders")
		return
	}
	// Sort folder
	if sort_by != nil && order_by != nil {
		// Sort by name
		if *sort_by == "sort_name" {
			if *order_by == "asc" {
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
			if *order_by == "asc" {
				sort.Slice(folders, func(i, j int) bool {
					return folders[i].Create_at > folders[j].Create_at
				})
			} else {
				sort.Slice(folders, func(i, j int) bool {
					return folders[i].Create_at < folders[j].Create_at
				})
			}
		}
	}
	// Transfer into string
	var folder_string []string
	for _, folder := range folders {
		folder_string = append(folder_string, fmt.Sprintf("%s|%s|%s|%s|%s", folder.ID, folder.Name, folder.Description, folder.Create_at, username))
	}
	// Print result
	for _, folder := range folder_string {
		fmt.Println(folder)
	}
	return
}

func Rename_folder(username string, folder_id string, new_folder_name string) {
	// Check user exist
	user, user_exist := Users[username]
	if !user_exist {
		fmt.Println("Error - unknown user")
		return
	}
	// Check specific folder exist
	folder, folder_exist := user.Folders[folder_id]
	if !folder_exist {
		fmt.Println("Error - folder_id not found")
		return
	}
	// rename the folder
	folder.Name = new_folder_name
	user.Folders[folder_id] = folder

	fmt.Println("Success")
	return
}

func Upload_file(username string, folder_id string, file_name string, description string) {
	// Check user exist
	user, user_exist := Users[username]
	if !user_exist {
		fmt.Println("Error - unknown user")
		return
	}
	// Check folder exist
	folder, folder_exist := user.Folders[folder_id]
	if !folder_exist {
		fmt.Println("Error - folder_id not found")
		return
	}
	// Create file
	file := File{
		File_name:   file_name,
		Extension:   strings.Split(file_name, ".")[1],
		Description: description,
		Create_at:   time.Now().Format("2006-01-02 15:04:05"),
	}
	// Put file in to folder
	folder.Files = append(folder.Files, file)
	user.Folders[folder_id] = folder

	fmt.Println("Success")
	return
}

func Delete_file(username string, folder_id string, file_name string) {
	// Check user exist
	user, user_exist := Users[username]
	if !user_exist {
		fmt.Println("Error - unknown user")
		return
	}
	// Checkt folder exist
	folder, folder_exist := user.Folders[folder_id]
	if !folder_exist {
		fmt.Println("Error - folder_id not found")
		return
	}
	// Select specific file and delete it
	for i, file := range folder.Files {
		if file.File_name == file_name {
			folder.Files = append(folder.Files[:i], folder.Files[i+1:]...)
			user.Folders[folder_id] = folder
			fmt.Println("Success")
			return
		}
	}
	// If didn't find out specific file
	fmt.Println("file_name not found")
	return
}

func Get_files(username string, folder_id string, sort_by *string, order_by *string) {
	// Check user exist
	user, user_exist := Users[username]
	if !user_exist {
		fmt.Println("Error - unknown user")
		return
	}
	// Check folder exist
	folder, folder_exist := user.Folders[folder_id]
	if !folder_exist {
		fmt.Println("Error - folder_id not found")
		return
	}
	// Check amount of files
	if len(folder.Files) == 0 {
		fmt.Println("Warning - empty files")
	}
	// Sort files
	if sort_by != nil && order_by != nil {
		// Sort by name
		if *sort_by == "sort_name" {
			if *order_by == "asc" {
				sort.Slice(folder.Files, func(i, j int) bool {
					return folder.Files[i].File_name > folder.Files[j].File_name
				})
			} else {
				sort.Slice(folder.Files, func(i, j int) bool {
					return folder.Files[i].File_name < folder.Files[j].File_name
				})
			}
			// Sort by time
		} else if *sort_by == "sort_time" {
			if *order_by == "asc" {
				sort.Slice(folder.Files, func(i, j int) bool {
					return folder.Files[i].Create_at > folder.Files[j].Create_at
				})
			} else {
				sort.Slice(folder.Files, func(i, j int) bool {
					return folder.Files[i].Create_at < folder.Files[j].Create_at
				})
			}
			// Sort by extension
		} else if *sort_by == "sort_extension" {
			if *order_by == "asc" {
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
	var files_string []string
	for _, file := range folder.Files {
		files_string = append(files_string, fmt.Sprintf("%s|%s|%s|%s|%s", file.File_name, file.Extension, file.Description, file.Create_at, username))
	}
	// Print the result out
	for _, file := range files_string {
		fmt.Println(file)
	}
	return
}

func main() {
	for true {
		fmt.Print("# ")
		// To get command-line input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		// Check methods
		function := strings.Split(input, " ")
		if function[0] == "register" {
			username := function[1]
			Register(username)
		} else if function[0] == "create_folder" {
			username := function[1]
			folder_name := strings.Trim(function[2], "‘’")
			description := strings.Join(function[3:], " ")
			description = strings.Trim(description, "‘’")
			Create_folder(username, folder_name, description)
		} else if function[0] == "delete_folder" {
			username := function[1]
			folder_id := function[2]
			Delete_folder(username, folder_id)
		} else if function[0] == "get_folders" {
			if len(function) > 2 {
				username := function[1]
				sort_by := function[2]
				order_by := function[3]
				Get_folders(username, &sort_by, &order_by)
				// if no specific sort way
			} else {
				username := function[1]
				Get_folders(username, nil, nil)
			}
		} else if function[0] == "rename_folder" {
			username := function[1]
			folder_id := function[2]
			folder_name := strings.Trim(function[3], "‘’")
			Rename_folder(username, folder_id, folder_name)
		} else if function[0] == "upload_file" {
			username := function[1]
			folder_id := function[2]
			file_name := strings.Trim(function[3], "‘’")
			description := strings.Join(function[4:], " ")
			description = strings.Trim(description, "‘’")
			Upload_file(username, folder_id, file_name, description)
		} else if function[0] == "delete_file" {
			username := function[1]
			folder_id := function[2]
			file_name := function[3]
			Delete_file(username, folder_id, file_name)
		} else if function[0] == "get_files" {
			if len(function) > 3 {
				username := function[1]
				folder_id := function[2]
				sort_by := function[3]
				order_by := function[4]
				Get_files(username, folder_id, &sort_by, &order_by)
				// if no specific sort way
			} else {
				username := function[1]
				folder_id := function[2]
				Get_files(username, folder_id, nil, nil)
			}
		} else {
			fmt.Println(Users)
		}
	}
}
