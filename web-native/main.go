package main

import (
	"fmt"
	"log"
	"net/http"
	"web-native/config"
	"web-native/controllers/categorycontroller"
	homecontrollers "web-native/controllers/homecontroller"
)

func main() {
	fmt.Println("Golang CRUD Web")
	config.ConnectDB()

	http.HandleFunc("/", homecontrollers.Welcome)

	http.HandleFunc("/category", categorycontroller.Index)
	http.HandleFunc("/category/add", categorycontroller.Add)
	http.HandleFunc("/category/edit", categorycontroller.Edit)
	http.HandleFunc("/categorycontroller", categorycontroller.Delete)

	log.Println("Server is running")
	http.ListenAndServe(":6060", nil)

}
