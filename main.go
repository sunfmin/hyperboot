package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sunfmin/hyperboot/services"
	"os"
)

func main() {

	var module string

	var hyperbootCmd = &cobra.Command{
		Use:   "hyperboot",
		Short: "Hyperboot is a application scafford generator",
		Long: `An application scafford generator prefering certain directory and module structure, like Ruby on Rails command.
Complete documentation is available at http://github.com/sunfmin/hyperboot`,
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hyperboot",
		Long:  `All software has versions. This is Hyperboot's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hyperboot v0.1 -- HEAD")
		},
	}

	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "Add initial configuration file",
		Long:  `To your modules config directory, add initial sample config file. If it already exists, skip creating.`,
		Run: func(cmd *cobra.Command, args []string) {
			pwd, _ := os.Getwd()
			err := services.CreateConfig(pwd, module, "mysql", false)
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	var dbCmd = &cobra.Command{
		Use:   "db",
		Short: "Add initial database file",
		Long:  `To your modules config directory, add initial sample database file. If it already exists, skip creating.`,
		Run: func(cmd *cobra.Command, args []string) {
			pwd, _ := os.Getwd()
			err := services.CreateDb(pwd, module)
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	hyperbootCmd.PersistentFlags().StringVarP(&module, "module", "m", "root", "module directory inside which to generate files.")
	hyperbootCmd.AddCommand(versionCmd)
	hyperbootCmd.AddCommand(configCmd)
	hyperbootCmd.AddCommand(dbCmd)

	hyperbootCmd.Execute()
}
