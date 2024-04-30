package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	cmd "koopa/cmd/db"
	"koopa/models"
	"log"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "db commands",
	Run: func(cmd *cobra.Command, args []string) {

		activeDatabase, err := models.LoadSetting("ACTIVE_DATABASE")
		if err != nil {
			log.Fatalf("error loading setting: %v", err)
		}

		message := fmt.Sprintf("You are using database: %s \n", activeDatabase.Value)
		fmt.Println(message)

		_ = cmd.Usage()

	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(cmd.NewDbCmd)
	dbCmd.AddCommand(cmd.ListDbCmd)
	dbCmd.AddCommand(cmd.CheckoutDbCmd)
}
