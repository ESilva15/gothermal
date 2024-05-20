package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"thermalFax/models"
)

type PrintServerResponse struct {
	Message string `json:"result"`
}

type WebFrontendRequest struct {
	Message string `json:"fax_message"`
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
		return
	}

	session, _ := models.GetSession(token)
	log.Printf(" | User %s request a fax.", session.User)

	data, _ := io.ReadAll(r.Body)
	log.Printf("| Received Request: %s", data)

	// Parse the JSON data that comes from the form
	var formData WebFrontendRequest
	err := json.Unmarshal(data, &formData)
	if err != nil {
		log.Println("Failed to unmarshal form data: ", err)
		w.Write([]byte("Failed to unmarshal form data."))
		return
	}

	resp, err := sendRequestToPrinter(formData.Message)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Failed to send message to printer: %s", err)))
		return
	}

	log.Println("Print Service Response: ", resp.Message)

	w.Write([]byte("Success"))
}
