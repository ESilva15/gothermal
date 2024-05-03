package server

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// pass -> $2y$08$ZBb/RxlrNHcA.YNasvGlhOeUmJVZe/clNelP4jwBpkpq/r4mE.lze
func isAuthorised(user string, pass string) bool {
	p := []byte("$2y$08$ZBb/RxlrNHcA.YNasvGlhOeUmJVZe/clNelP4jwBpkpq/r4mE.lze")
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
