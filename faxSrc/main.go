package main

import (
	"fmt"
	"log"
	"net/http"
	"thermalFax/controllers"
	"thermalFax/models"
	"thermalFax/views"
)

// launchServer is the main server function
func launchServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	// Authenticate the user
	// TODO: I should make this a bit better since we want to use another service
	// and that one should handle the authentication and this one make use of it
	var auth models.Authentication
	if !auth.Authenticate(w, r) {
		return
	}

	views.EditorPage(w, r)
}

func main() {
	// Set up a very basic http server
	http.HandleFunc("/", launchServer)

	http.HandleFunc("/send-fax", controllers.SendFax)

	// Serve the CSS and JS files, we might have to change this
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Starting server on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
