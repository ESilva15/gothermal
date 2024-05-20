package controllers

import (
	"net/http"
	"thermalFax/views"
)

func EditorPage(w http.ResponseWriter, r *http.Request) {
	authd, _ := isAuthenticated(r)
	if !authd {
		views.LoginPage(w, r)
		return
	}

	views.EditorPage(w, r)
}
