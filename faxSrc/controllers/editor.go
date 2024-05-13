package controllers

import (
	"net/http"
	"thermalFax/views"
)

func EditorPage(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		views.LoginPage(w, r)
		return
	}

	views.EditorPage(w, r)
}
