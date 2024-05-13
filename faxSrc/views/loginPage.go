package views

import "net/http"

func LoginPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	loginPage := htmlTemplate("./assets/htmx/terminal.tmpl", map[string]interface{}{
		"content":      htmlTemplate("./assets/htmx/login.html", nil),
		"renderNavBar": true,
		"page":         "fax",
	})

	_, err := w.Write([]byte(loginPage))
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
