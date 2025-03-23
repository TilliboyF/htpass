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

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all users",
	Long:    "List all users",
	Args:    cobra.NoArgs,
	Aliases: []string{"ls"},
	Run:     handleList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func handleList(cmd *cobra.Command, args []string) {
	users, err := common.ReadHtpasswdFile(htPasswdFile)
	if err != nil {
		log.Fatal(err)
	}

	headerFmt := color.New(color.FgCyan, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("users", "hashed passwords")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for user, pw := range users {
		tbl.AddRow(user, pw)
	}

	buf := new(bytes.Buffer)
	tbl.WithWriter(buf)
	tbl.Print()
	fmt.Println(buf.String())
}
