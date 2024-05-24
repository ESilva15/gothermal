package views

import "net/http"

func EditorPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	mainPage := HtmlTemplate("./assets/htmx/terminal.tmpl", map[string]interface{}{
		"content":      HtmlTemplate("./assets/htmx/editor.html", nil),
		"renderNavBar": true,
		"page":         "fax",
	})

	_, err := w.Write([]byte(mainPage))
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
