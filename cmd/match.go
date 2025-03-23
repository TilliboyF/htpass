package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tilliboyf/htpass/common"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"log"
	"os"
)

var matchCmd = &cobra.Command{
	Use:     "match [username]",
	Short:   "Matches a password of a user",
	Long:    `Matches a password of a user`,
	Args:    cobra.ExactArgs(1),
	Run:     handleMatch,
	Aliases: []string{"m"},
}

func handleMatch(cmd *cobra.Command, args []string) {
	users, err := common.ReadHtpasswdFile(htPasswdFile)
	if err != nil {
		log.Fatal(err)
	}
	username := args[0]
	if hashedPw, ok := users[username]; ok {
		fmt.Printf("Enter password to match for user %s: ", username)
		enteredPw, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Printf("\n")
		if err != nil {
			log.Fatalf("Error reading password: %v", err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPw), enteredPw)
		if err != nil {
			fmt.Println(color.RedString("Passwords do not match"))
		} else {
			fmt.Println(color.GreenString("Password match"))
		}

	} else {
		log.Fatalf("User %s not found", username)
	}
}

func init() {
	rootCmd.AddCommand(matchCmd)
}
