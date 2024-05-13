package views

import "net/http"

func EditorPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	mainPage := htmlTemplate("./assets/htmx/terminal.tmpl", map[string]interface{}{
		"content":      htmlTemplate("./assets/htmx/editor.html", nil),
		"renderNavBar": true,
		"page":         "fax",
	})

	_, err := w.Write([]byte(mainPage))
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
