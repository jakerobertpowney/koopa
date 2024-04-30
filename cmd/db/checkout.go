package cmd

import (
	"github.com/spf13/cobra"
	"koopa/models"
	"log"
)

var Label string

var CheckoutDbCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout a models to another models",
	Run: func(cmd *cobra.Command, args []string) {

		_, err := models.LoadDatabase(Label)
		if err != nil {
			log.Fatalf("Error loading database: %v", err)
		}

		setting := models.Setting{
			Value: Label,
		}

		if err := setting.Save("ACTIVE_DATABASE"); err != nil {
			log.Fatalf("Error saving setting: %v", err)
		}

	},
}

func init() {
	CheckoutDbCmd.Flags().StringVarP(&Label, "label", "l", "", "Database name that you want to connect to")
	_ = CheckoutDbCmd.MarkFlagRequired("label")
}
