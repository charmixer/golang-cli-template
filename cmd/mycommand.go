package cmd

import (
	"github.com/charmixer/golang-cli-template/pkg/mycmd"
	"github.com/spf13/cobra"
)

var mycommand = &cobra.Command{
	Use:   "mycommand",
	Short: "mycommand",
	Long:  `mycommand...`,
	Run:   mycmd.RunMyCommand(),
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mycommand.PersistentFlags().String("foo", "", "A help for foo")

	//mycommand.Flags().StringP("flag", "f", "default", "Some flag needed by mycommand")

	rootCmd.AddCommand(mycommand)
}
