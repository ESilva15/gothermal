/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	thermalprinter "thermalPrinter/thermalPrinter"

	"github.com/spf13/cobra"
)

func messageCmdAction(cmd *cobra.Command, args []string) {
	message, _ := cmd.Flags().GetString("message")

	printer := thermalprinter.GetInstance(nil)

	printer.Print(message)
}

// messageCmd represents the message command
var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "",
	Long:  ``,
	Run:   messageCmdAction,
}

func init() {
	rootCmd.AddCommand(messageCmd)

	messageCmd.Flags().StringP("message", "m", "Hello, world!\n", "Message to print")
	messageCmd.MarkFlagRequired("message")
}
