package models

// This isn't used anymore, will keep it here for a while, for reasons, its mine

// pass -> $2y$08$ZBb/RxlrNHcA.YNasvGlhOeUmJVZe/clNelP4jwBpkpq/r4mE.lze
// func (a *Authentication) isAuthorised(user string, pass string) bool {
// 	p, ok := DATABASE[user]
// 	if !ok {
// 		return false
// 	}

// 	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(pass))

// 	return err == nil
// }

// func (a *Authentication) Authenticate(w http.ResponseWriter, r *http.Request) bool {
// 	username, password, ok := r.BasicAuth()
// 	if !ok {
// 		log.Println("Log in failed")
// 		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
// 		w.WriteHeader(http.StatusUnauthorized)
// 		w.Write([]byte(`{"message": "No basic auth present"}`))
// 		return false
// 	}

// 	if !a.isAuthorised(username, password) {
// 		log.Printf("Log in failed: %s:%s", username, password)
// 		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
// 		w.WriteHeader(http.StatusUnauthorized)
// 		w.Write([]byte(`{"message": "Invalid username or password"}`))
// 		return false
// 	}

// 	log.Printf("User %s logged in successfully", username)
// 	return true
// }
