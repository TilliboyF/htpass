package cmd

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/tilliboyf/htpass/common"
	"log"
)

var getCmd = &cobra.Command{
	Use:   "get [username]",
	Short: "get a user",
	Long:  "get a user",
	Args:  cobra.ExactArgs(1),
	Run:   handleGet,
}

func handleGet(cmd *cobra.Command, args []string) {
	users, err := common.ReadHtpasswdFile(htPasswdFile)
	if err != nil {
		log.Fatal(err)
	}
	username := args[0]
	if hashedPw, ok := users[username]; ok {

		headerFmt := color.New(color.FgCyan, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("users", "hashed passwords")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		tbl.AddRow(username, hashedPw)

		buf := new(bytes.Buffer)
		tbl.WithWriter(buf)
		tbl.Print()
		fmt.Println(buf.String())
	} else {
		fmt.Printf("User %s not found\n", username)
	}
}

func init() {
	rootCmd.AddCommand(getCmd)
}
