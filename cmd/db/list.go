package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"koopa/models"
	"log"
)

var ListDbCmd = &cobra.Command{
	Use:   "list",
	Short: "List all databases",
	Run: func(cmd *cobra.Command, args []string) {

		databases, err := models.LoadAllDatabases()
		if err != nil {
			log.Fatalf("error loading databases: %v", err)
		}

		fmt.Println("Databases:")

		for _, db := range databases {
			database := fmt.Sprintf("%s - %s", db.Database, db.Type)
			fmt.Println(database)
		}

	},
}
