/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	thermalprinter "thermalPrinter/thermalPrinter"

	"github.com/spf13/cobra"
)

func rootCmdAction(cmd *cobra.Command, args []string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")

	// Scan until EOF (Ctrl+D)
	text := ""
	// Should refactor this weird looking thing
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	printer := thermalprinter.GetInstance(nil)
	printer.Print(text)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thermalPrinter",
	Short: "",
	Long:  ``,
	Run:   rootCmdAction,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	/* 	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.thermalPrinter.yaml)") */

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	/* 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle") */
}
