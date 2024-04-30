package cmd

import (
	"github.com/spf13/cobra"
	"koopa/models"
	"log"
)

var Label string
var User string
var Hostname string
var Port string

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new SSH connection",
	Run: func(cmd *cobra.Command, args []string) {

		connection := models.Connection{
			Label:    Label,
			Hostname: Hostname,
			Port:     Port,
			Username: User,
		}

		err := connection.Save()
		if err != nil {
			log.Fatalf("Error saving connection: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&Label, "label", "l", "", "SSH Label (You will be use this to connect to the server)")
	_ = newCmd.MarkFlagRequired("label")
	newCmd.Flags().StringVarP(&User, "user", "u", "", "SSH Username")
	_ = newCmd.MarkFlagRequired("user")
	newCmd.Flags().StringVarP(&Hostname, "host", "i", "", "SSH Host Name")
	_ = newCmd.MarkFlagRequired("host")
	newCmd.Flags().StringVarP(&Port, "port", "p", "", "SSH Port")
	_ = newCmd.MarkFlagRequired("port")
}
