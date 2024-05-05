package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// pass -> $2y$08$ZBb/RxlrNHcA.YNasvGlhOeUmJVZe/clNelP4jwBpkpq/r4mE.lze
func isAuthorised(user string, pass string) bool {
	p := []byte("$2y$08$ZBb/RxlrNHcA.YNasvGlhOeUmJVZe/clNelP4jwBpkpq/r4mE.lze")

	if user != "r" {
		return false
	}

	err := bcrypt.CompareHashAndPassword(p, []byte(pass))

	return err == nil
}

func authenticate(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "No basic auth present"}`))
		return false
	}

	if !isAuthorised(username, password) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid username or password"}`))
		return false
	}

	return true
}

func launchServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	if !authenticate(w, r) {
		return
	}

	// data, _ := io.ReadAll(r.Body)

	// main page starts here I guess
	w.WriteHeader(http.StatusOK)

	mainPage := HtmlTemplate("./assets/htmx/terminal.tmpl", map[string]interface{}{
		"content":      HtmlTemplate("./assets/htmx/login.html", nil),
		"renderNavBar": true,
		"page":         "fax",
	})

	_, err := w.Write([]byte(mainPage))
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Set up a very basic http server
	http.HandleFunc("/", launchServer)

	// Serve the CSS and JS files, we might have to change this
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Starting server on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
