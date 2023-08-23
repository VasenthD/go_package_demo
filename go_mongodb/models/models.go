package models

type Stuents struct{
	Title        string             `json:"title" bson:"title"`
	Author       string             `json:"author" bson:"author"`
	PublishedYear int                `json:"publishedYear" bson:"publishedYear"`
}