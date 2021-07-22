package models

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Abstract string `json:"abstract"`
}
