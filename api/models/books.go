package models

type Book struct {
	Title string `json:"title"`
	Abstract string `json:"abstract"`
	Author string `json:"author"`
}