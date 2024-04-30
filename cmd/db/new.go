package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"koopa/models"
	"log"
	_ "modernc.org/sqlite"
)

var NewDbCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new db",
	Run: func(cmd *cobra.Command, args []string) {

		typePrompt := promptui.Select{
			Label: "Select Database Host",
			Items: []string{"Local", "MySQL (Remote)"},
		}

		_, dbType, err := typePrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if dbType == "Local" {

			labelPrompt := promptui.Prompt{
				Label: "Database Name",
			}

			dbName, err := labelPrompt.Run()
			if err != nil {
				log.Fatalf("Label prompt failed %v", err)
			}

			database := models.Database{
				Type:     "local",
				Database: dbName,
				User:     "",
				Password: "",
				Hostname: "",
				Port:     "",
			}

			dbFilename := fmt.Sprintf("./%s.db", database.Database)
			db, err := sql.Open("sqlite", dbFilename)

			if err != nil {
				log.Fatalf("Failed to open database: %v", err)
			}

			statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS connections (label TEXT PRIMARY KEY, username TEXT, hostname TEXT, port TEXT)")
			if _, err := statement.Exec(); err != nil {
				log.Fatalf("Failed to create connections table: %v", err)
			}

			if err := database.Save(database.Database); err != nil {
				log.Fatalf("Error saving database: %v", err)
			}

		} else if dbType == "MySQL (Remote)" {

			database := models.Database{
				Type: "MySQL (Remote)",
			}

			prompt := promptui.Prompt{
				Label: "Database Name",
			}
			database.Database, err = prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
			}

			prompt = promptui.Prompt{
				Label: "Username",
			}
			database.User, err = prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
			}

			prompt = promptui.Prompt{
				Label: "Password",
			}
			database.Password, err = prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
			}

			prompt = promptui.Prompt{
				Label: "Hostname (IP Address)",
			}
			database.Hostname, err = prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
			}

			prompt = promptui.Prompt{
				Label: "Port",
			}
			database.Port, err = prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
			}

			if err := database.Save(database.Database); err != nil {
				log.Fatalf("Error saving database: %v", err)
			}

		}

	},
}
