package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tilliboyf/htpass/common"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"log"
	"os"
)

var bcryptCost int

var setPasswordCmd = &cobra.Command{
	Use:     "set-password [username]",
	Short:   "Set a user's password (add or update)",
	Long:    "Set a user's password (add or update)",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"set", "sp"},
	Run:     handleSetPassword,
}

func init() {
	setPasswordCmd.Flags().IntVarP(&bcryptCost, "cost", "c", bcrypt.DefaultCost, "Bcrypt cost")
	rootCmd.AddCommand(setPasswordCmd)
}

func handleSetPassword(cmd *cobra.Command, args []string) {
	users, err := common.ReadHtpasswdFile(htPasswdFile)
	if err != nil {
		log.Fatal(err)
	}
	username := args[0]
	fmt.Printf("Enter new password for user %s: ", username)
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("Error reading password: %v", err)
	}
	password := string(bytePassword)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	users[username] = string(hashedPassword)

	err = common.WriteHtpasswdFile(htPasswdFile, users)
	if err != nil {
		log.Fatal(err)
	}

	for i := range bytePassword {
		bytePassword[i] = 0
	}
}
