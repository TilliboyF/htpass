package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tilliboyf/htpass/common"
	"log"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [username]",
	Short:   "Delete user",
	Long:    "Delete user",
	Aliases: []string{"del"},
	Args:    cobra.ExactArgs(1),
	Run:     handleDelete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func handleDelete(cmd *cobra.Command, args []string) {
	username := args[0]
	users, err := common.ReadHtpasswdFile(htPasswdFile)
	if err != nil {
		log.Fatal(err)
	}
	if val := users[username]; val == "" {
		fmt.Println(color.RedString("user %s not found", username))
	} else {
		delete(users, username)
		err = common.WriteHtpasswdFile(htPasswdFile, users)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(color.GreenString("user %s deleted", username))
	}
}
