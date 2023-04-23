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
	Label       []string
	Files       []File
}

// Label struct
type Label struct {
	Label_name   string
	Color        string
	Created_time string
}

// User structure
type User struct {
	Labels  map[string]Label
	Folders map[string]Folder
}

// Create a Users object
var Users = make(map[string]User)

// Count folder ID
var Id_count int = 1001

func Check_user(username string) (User, error) {
	user, user_exist := Users[username]
	if !user_exist {
		return User{}, fmt.Errorf("Error - unknown user")
	}
	return user, nil
}

func Check_folder(user User, folder_id string) (Folder, error) {
	folder, folder_exist := user.Folders[folder_id]
	if !folder_exist {
		return Folder{}, fmt.Errorf("Error - folder_id not found")
	}
	return folder, nil
}

func Register(username string) {
	// Check user exist
	if _, user_exist := Users[username]; user_exist {
		fmt.Println("Error - user already existing")
		return
	}
	// Create User
	Users[username] = User{
		Folders: make(map[string]Folder),
		Labels:  make(map[string]Label),
	}

	fmt.Println("Success")
	return
}

func Create_folder(username string, folder_name string, description string) {
	// Check user exist
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
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
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
		return
	}
	// Check folder exist
	_, folder_err := Check_folder(user, folder_id)
	if folder_err != nil {
		fmt.Println("Error - folder doesn't exist")
		return
	}
	// Delete folder
	delete(user.Folders, folder_id)
	fmt.Println("Success")
	return
}

func Get_folders(username string, label_name *string, sort_by *string, order_by *string) {
	var folders []Folder
	// Check user exist
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
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
	// Check label exist
	if label_name != nil {
		if _, label_exist := user.Labels[*label_name]; !label_exist {
			fmt.Println("Error - label name not exist")
			return
		}
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
		if label_name != nil {
			folder_string = append(folder_string, fmt.Sprintf("%s|%s|%s|%s|%s|%s", folder.ID, *label_name, folder.Name, folder.Description, folder.Create_at, username))
		} else {
			folder_string = append(folder_string, fmt.Sprintf("%s|%s|%s|%s|%s", folder.ID, folder.Name, folder.Description, folder.Create_at, username))
		}
	}
	// Print result
	for _, folder := range folder_string {
		fmt.Println(folder)
	}
	return
}

func Rename_folder(username string, folder_id string, new_folder_name string) {
	// Check user exist
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
		return
	}
	// Check specific folder exist
	folder, folder_err := Check_folder(user, folder_id)
	if folder_err != nil {
		fmt.Println(folder_err)
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
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
		return
	}
	// Check folder exist
	folder, folder_err := Check_folder(user, folder_id)
	if folder_err != nil {
		fmt.Println(folder_err)
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
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
	}
	// Check folder exist
	folder, folder_err := Check_folder(user, folder_id)
	if folder_err != nil {
		fmt.Println(folder_err)
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
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
		return
	}
	// Check folder exist
	folder, folder_err := Check_folder(user, folder_id)
	if folder_err != nil {
		fmt.Println(folder_err)
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

func Add_label(username string, label_name string, color string) {
	// Check user exist
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
		return
	}
	// Check label already exist
	for name, _ := range user.Labels {
		if name == label_name {
			fmt.Println("the label name existing")
			return
		}
	}
	// Create label and add in user
	label := Label{
		Color:        color,
		Created_time: time.Now().Format("2006-01-02 15:04:05"),
	}
	user.Labels[label_name] = label
	Users[username] = user

	fmt.Println("Success")
}

func Get_labels(username string) {
	var labels []Label
	// Check user exist
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
		return
	}
	// Append all label into labels
	for label_name, label := range user.Labels {
		label.Label_name = label_name
		labels = append(labels, label)
	}
	// if no label
	if len(labels) == 0 {
		fmt.Println("Warning - empty labels")
	}
	// Transfer to string
	var label_string []string
	for _, label := range labels {
		label_string = append(label_string, fmt.Sprintf("%s|%s|%s|%s", label.Label_name, label.Color, label.Created_time, username))
	}
	// Print the result
	for _, label := range label_string {
		fmt.Println(label)
	}
	return
}

func Delete_labels(username string, label_name string) {
	// Check user exist
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
		return
	}
	// Check label exist
	_, label_exist := user.Labels[label_name]
	if !label_exist {
		fmt.Println("Error - the label name not exist")
		return
	}
	// Delete label
	delete(user.Labels, label_name)
	fmt.Println("Success")
	return
}

func Add_folder_label(username string, folder_id string, label_name string) {
	// Check user exist
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
		return
	}
	// Check folder exist
	folder, folder_err := Check_folder(user, folder_id)
	if folder_err != nil {
		fmt.Println(folder_err)
		return
	}
	// Check label exist
	for _, label := range folder.Label {
		if label == label_name {
			fmt.Println("Error - label name already exist")
			return
		}
	}
	folder.Label = append(folder.Label, label_name)
	fmt.Println("Success")
	return
}

func Delete_folder_label(username string, folder_id string, label_name string) {
	// Check user exist
	user, user_err := Check_user(username)
	if user_err != nil {
		fmt.Println(user_err)
		return
	}
	// Check folder exist
	folder, folder_err := Check_folder(user, folder_id)
	if folder_err != nil {
		fmt.Println(folder_err)
		return
	}
	// Delete folder label
	for i, label := range folder.Label {
		if label == label_name {
			folder.Label = append(folder.Label[:i], folder.Label[i+1:]...)
			user.Folders[folder_id] = folder
			fmt.Println("Success")
			return
		}
	}
	// If didn't find out specific file
	fmt.Println("label not found")
	return
}

func main() {
	for true {
		fmt.Print("# ")
		// To get command-line input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := string(scanner.Text())
		// read the input into key word
		var word_string []string
		var quote bool = false
		var word string
		for _, letter := range input {
			if !quote {
				if letter == ' ' {
					word_string = append(word_string, word)
					word = ""
				} else if letter == '\'' {
					quote = true
				} else {
					word = word + string(letter)
				}
			} else {
				if letter == '\'' {
					quote = false
					word_string = append(word_string, word)
					word = ""
				} else {
					word = word + string(letter)
				}
			}
		}
		if input[len(input)-1] != '\'' || input[len(input)-1] != ' ' {
			word_string = append(word_string, word)
		}
		// move item which is not "" to function
		var function []string
		for _, file := range word_string {
			word = strings.Trim(file, " ")
			if word != "" {
				function = append(function, word)
			}
		}
		// Check methods
		if function[0] == "register" {
			Register(function[1])
		} else if function[0] == "create_folder" {
			Create_folder(function[1], function[2], function[3])
		} else if function[0] == "delete_folder" {
			Delete_folder(function[1], function[2])
		} else if function[0] == "get_folders" {
			if len(function) > 2 {
				Get_folders(function[1], &function[2], &function[3], &function[4])
			} else {
				// if no specific sort way
				Get_folders(function[1], nil, nil, nil)
			}
		} else if function[0] == "rename_folder" {
			Rename_folder(function[1], function[2], function[3])
		} else if function[0] == "upload_file" {
			Upload_file(function[1], function[2], function[3], function[4])
		} else if function[0] == "delete_file" {
			Delete_file(function[1], function[2], function[3])
		} else if function[0] == "get_files" {
			if len(function) > 3 {
				Get_files(function[1], function[2], &function[3], &function[4])
			} else {
				// if no specific sort way
				Get_files(function[1], function[2], nil, nil)
			}
		} else if function[0] == "add_label" {
			Add_label(function[1], function[2], function[3])
		} else if function[0] == "get_labels" {
			Get_labels(function[1])
		} else if function[0] == "delete_labels" {
			Delete_labels(function[1], function[2])
		} else if function[0] == "add_folder_label" {
			Add_folder_label(function[1], function[2], function[3])
		} else if function[0] == "delete_folder_label" {
			Delete_folder_label(function[1], function[2], function[3])
		}
	}
}
