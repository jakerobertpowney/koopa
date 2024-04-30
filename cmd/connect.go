package cmd

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"koopa/models"
	"log"
	_ "modernc.org/sqlite"
	"os"
	"os/exec"
)

var (
	Connection string
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a server using SSH.",
	Run: func(cmd *cobra.Command, args []string) {

		message := fmt.Sprintf("Trying to connect to %s", Connection)
		fmt.Println(message)

		connection, err := models.LoadConnection("staging.its")
		if err != nil {
			log.Fatalf("Error loading connection: %s", err)
		}

		connectionString := fmt.Sprintf("%s@%s", connection.Username, connection.Hostname)
		connectionPort := fmt.Sprintf("-p %s", connection.Port)

		command := exec.Command("ssh", connectionString, connectionPort)
		// redirect the output to terminal
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Stdin = os.Stdin

		err = command.Run()
		if err != nil {
			log.Fatalf("Error connecting to server: %s", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringVarP(&Connection, "connection", "c", "", "Server to connect to")
}
