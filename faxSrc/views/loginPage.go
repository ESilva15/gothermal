package views

import "net/http"

func LoginPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	loginTemplate := HtmlTemplate("./assets/htmx/login.html", map[string]interface{}{
		"script": true,
	})

	loginPage := HtmlTemplate("./assets/htmx/terminal.tmpl", map[string]interface{}{
		"content":      loginTemplate,
		"renderNavBar": true,
		"page":         "fax",
	})

	_, err := w.Write([]byte(loginPage))
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
