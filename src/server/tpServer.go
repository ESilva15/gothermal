package server

import (
	"fmt"
	"io"
	"net/http"
	thermalprinter "thermalPrinter/thermalPrinter"
)

func printMessage(message string) {
	printer := thermalprinter.GetInstance(nil)
	printer.Print(message)
}

func LaunchServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	data, _ := io.ReadAll(r.Body)
	fmt.Println(string(data))
	printMessage(string(data) + "\n\n\n")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result": "SUCCESS"}`))
}
