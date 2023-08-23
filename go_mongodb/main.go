package main

import (
	"fmt"
	"mongo/connections"
)

func main() {
	fmt.Println("hi")
	connections.Init()

	// connections.Inserts(&models.Stuents{Title: "alse Book",
	// 	Author:        "Jkdkdke",
	// 	PublishedYear: 8800})
	//connections.Updates("64e4dd2059a1ebe4c7f1bd42")
	connections.Deleteone("64e4dd2059a1ebe4c7f1bd42")

}
