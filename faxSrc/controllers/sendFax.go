package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

	"thermalFax/models"
)

type PrintServerResponse struct {
	Message string `json:"result"`
}

type WebFrontendRequest struct {
	Message string `json:"fax_message"`
}

type WebFrontendResponse struct {
	State string `json:"state"`
}

// SendRequestToPrinter function will request the thermal printer
// to print the message.
// It should be able to send the requests both via http/https or sockets
func sendRequestToPrinter(msg string) (PrintServerResponse, error) {
	// Send the request to the actual printer server
	req, err := http.NewRequest(
		http.MethodPost,
		"http://127.0.0.1:9099/print",
		bytes.NewBuffer([]byte(msg)),
	)
	if err != nil {
		errStr := fmt.Sprintf("Failed to create new request: %s", err)
		return PrintServerResponse{}, errors.New(errStr)
	}

	// Perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		errStr := fmt.Sprintf("Failed to perform new request: %s", err)
		return PrintServerResponse{}, errors.New(errStr)
	}

	// Process the request data
	out, _ := io.ReadAll(res.Body)
	var resp PrintServerResponse
	err = json.Unmarshal(out, &resp)
	if err != nil {
		errStr := fmt.Sprintf("Failed to unmarshall response: %s", err)
		return PrintServerResponse{}, errors.New(errStr)
	}

	return resp, nil
}

// sendFax is a function that receives the data from the fronted and
// sends it to the print server
func SendFax(w http.ResponseWriter, r *http.Request) {
	authd, token := isAuthenticated(r)
	if !authd {
		w.Header().Set("Content-Type", "text/json")
		w.Write([]byte("{\"state\":\"unauthenticated\"}"))

		return
	}

	session, _ := models.GetSession(token)
	log.Printf("| User %s requested a fax", session.User)

	data, _ := io.ReadAll(r.Body)

	// Parse the JSON data that comes from the form
	var formData WebFrontendRequest
	err := json.Unmarshal(data, &formData)
	if err != nil {
		log.Println("Failed to unmarshal form data: ", err)
		w.Write([]byte("{\"state\":\"failed\"}"))
		return
	}

	msg := signMessage(formData.Message, session.User, time.Now().Format("2006-01-02 15:04:05"))
	_, err = sendRequestToPrinter(msg)
	if err != nil {
		log.Println("| Printer failed to process fax:", err)
		w.Write([]byte("{\"state\":\"failed\"}"))
		return
	}

	w.Header().Set("Content-Type", "text/json")
	w.Write([]byte("{\"state\":\"success\"}"))
}

// NOTE: this isn't multi-Unicode-codepoint aware, like specifying skintone or
// gender of an emoji: https://unicode.org/emoji/charts/full-emoji-modifiers.html
// This thing was straight up copied from SO
func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

// truncateString given a string and a lim truncates that string if its longer than the lim
// with "..."
func truncateString(s string, lim int) string {
	runeCount := utf8.RuneCountInString(s)
	if runeCount > lim {
		s = substr(s, 0, lim-3) + "..."
	} else {
		s = s + strings.Repeat(" ", lim-runeCount)
	}

	return s
}

// signMessage adds a small box to the message so the reader knows who sent it and the time
func signMessage(msg string, user string, time string) string {
	// defWidht := 48
	nameWidth := 39
	userLine := "| User: "
	userLine = userLine + truncateString(user, nameWidth) + "|"

	// defWidht := 48
	timeWidth := 39
	timeLine := "| Time: "
	timeLine = timeLine + truncateString(time, timeWidth) + "|"

	sign := fmt.Sprintf(`
+----------------------------------------------+
%s
%s
+----------------------------------------------+`,
		userLine, timeLine)
	truncation := "+- END ---------------------------------- END -+"

	return sign + "\n" + msg + "\n\n" + truncation
}
