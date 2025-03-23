package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var createFileCmd = &cobra.Command{
	Use:     "createFile",
	Short:   "Create htpasswd file",
	Long:    "Create htpasswd file",
	Args:    cobra.NoArgs,
	Aliases: []string{"cf"},
	Run:     handleCreateFile,
}

func init() {
	rootCmd.AddCommand(createFileCmd)
}

func handleCreateFile(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(htPasswdFile); os.IsNotExist(err) {
		file, err := os.Create(htPasswdFile)
		if err != nil {
			fmt.Printf("Failed to create file: %v\n", err)
			return
		}
		defer file.Close()
		fmt.Println("File created successfully.")
	} else if err != nil {
		fmt.Printf("Error checking file: %v\n", err)
	} else {
		fmt.Println("File already exists.")
	}
}
