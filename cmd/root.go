package cmd

import "github.com/spf13/cobra"

var htPasswdFile string

var rootCmd = &cobra.Command{
	Use:   "htpass",
	Short: "User password handler for htpasswd",
	Long:  "User password handler for htpasswd",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&htPasswdFile, "file", "f", ".htpasswd", "Location of your htpasswd file")
}

func Execute() error {
	return rootCmd.Execute()
}
