/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	thermalprinter "thermalPrinter/thermalPrinter"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

// pass -> $2y$08$ZBb/RxlrNHcA.YNasvGlhOeUmJVZe/clNelP4jwBpkpq/r4mE.lze
func isAuthorised(user string, pass string) bool {
	p := []byte("$2y$08$ZBb/RxlrNHcA.YNasvGlhOeUmJVZe/clNelP4jwBpkpq/r4mE.lze")
	err := bcrypt.CompareHashAndPassword(p, []byte(pass))

	return err == nil
}

func printMessage(message string) {
	printer := thermalprinter.GetInstance(nil)
	printer.Print(message)
}

func print(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "No basic auth present"}`))
		return
	}

	if !isAuthorised(username, password) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid username or password"}`))
		return
	}

	data, _ := io.ReadAll(r.Body)
	printMessage(string(data) + "\n\n\n\n\n")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "welcome to thermics"}`))
}

func serverCmdAction(cmd *cobra.Command, args []string) {
	http.HandleFunc("/print", print)
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	Run:   serverCmdAction,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
