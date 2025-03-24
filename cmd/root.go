package cmd

import "github.com/spf13/cobra"

var htPasswdFile string

var rootCmd = &cobra.Command{
	Use:     "htpass",
	Short:   "User password handler for htpasswd",
	Long:    "User password handler for htpasswd, that uses bcrypt under the hood and gives useful functionality to handle htpasswd management",
	Version: "1.0.0",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&htPasswdFile, "file", "f", ".htpasswd", "Location of your htpasswd file")
}

func Execute() error {
	return rootCmd.Execute()
}
