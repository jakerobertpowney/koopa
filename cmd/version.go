/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Koopa",
	Long:  `The current version of Koopa that you have installed locally.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Koopa v0.2.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
