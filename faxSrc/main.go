package main

import (
	"fmt"
	"log"
	"net/http"
	"thermalFax/controllers"
)

func checkMethod(method string, r *http.Request) bool {
	return r.Method == method
}

func main() {
	// GET calls
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !checkMethod(http.MethodGet, r) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		controllers.EditorPage(w, r)
	})

	// POST calls
	http.HandleFunc("/send-fax", func(w http.ResponseWriter, r *http.Request) {
		if !checkMethod(http.MethodPost, r) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		controllers.SendFax(w, r)
	})

	http.HandleFunc("/send-login", func(w http.ResponseWriter, r *http.Request) {
		if !checkMethod(http.MethodPost, r) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		controllers.Authenticate(w, r)
	})

	// Serve the CSS and JS files, we might have to change this
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Starting server on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
