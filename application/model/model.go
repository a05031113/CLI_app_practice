package model

// File structure
type File struct {
	FileName    string
	Extension   string
	Description string
	CreateAt    string
}

// Folder structure
type Folder struct {
	ID          string
	Name        string
	Description string
	CreateAt    string
	Label       []string
	Files       []File
}

// Label struct
type Label struct {
	LabelName   string
	Color       string
	CreatedTime string
}

// User structure
type User struct {
	Labels  map[string]Label
	Folders map[string]Folder
}
