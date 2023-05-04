package main

import (
	"bufio"
	"fmt"
	"iscool/application/controller"
	"os"
	"strings"
)

func main() {
	for true {
		fmt.Print("# ")
		// To get command-line input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := string(scanner.Text())
		// read the input into key word
		var wordString []string
		var quote bool = false
		var word string
		for _, letter := range input {
			if !quote {
				if letter == ' ' {
					wordString = append(wordString, word)
					word = ""
				} else if letter == '\'' {
					quote = true
				} else {
					word = word + string(letter)
				}
			} else {
				if letter == '\'' {
					quote = false
					wordString = append(wordString, word)
					word = ""
				} else {
					word = word + string(letter)
				}
			}
		}
		if input[len(input)-1] != '\'' || input[len(input)-1] != ' ' {
			wordString = append(wordString, word)
		}
		// move item which is not "" to function
		var function []string
		for _, file := range wordString {
			word = strings.Trim(file, " ")
			if word != "" {
				function = append(function, word)
			}
		}
		// Check methods
		if function[0] == "register" {
			fmt.Println(controller.Register(function[1]))
		} else if function[0] == "create_folder" {
			fmt.Println(controller.Create_folder(function[1], function[2], function[3]))
		} else if function[0] == "delete_folder" {
			fmt.Println(controller.Delete_folder(function[1], function[2]))
		} else if function[0] == "get_folders" {
			if len(function) > 2 {
				fmt.Println(controller.Get_folders(function[1], &function[2], &function[3], &function[4]))
			} else {
				// if no specific sort way
				fmt.Println(controller.Get_folders(function[1], nil, nil, nil))
			}
		} else if function[0] == "rename_folder" {
			fmt.Println(controller.Rename_folder(function[1], function[2], function[3]))
		} else if function[0] == "upload_file" {
			fmt.Println(controller.Upload_file(function[1], function[2], function[3], function[4]))
		} else if function[0] == "delete_file" {
			fmt.Println(controller.Delete_file(function[1], function[2], function[3]))
		} else if function[0] == "get_files" {
			if len(function) > 3 {
				fmt.Println(controller.Get_files(function[1], function[2], &function[3], &function[4]))
			} else {
				// if no specific sort way
				fmt.Println(controller.Get_files(function[1], function[2], nil, nil))
			}
		} else if function[0] == "add_label" {
			fmt.Println(controller.Add_label(function[1], function[2], function[3]))
		} else if function[0] == "get_labels" {
			fmt.Println(controller.Get_labels(function[1]))
		} else if function[0] == "delete_labels" {
			fmt.Println(controller.Delete_labels(function[1], function[2]))
		} else if function[0] == "add_folder_label" {
			fmt.Println(controller.Add_folder_label(function[1], function[2], function[3]))
		} else if function[0] == "delete_folder_label" {
			fmt.Println(controller.Delete_folder_label(function[1], function[2], function[3]))
		}
	}
}
