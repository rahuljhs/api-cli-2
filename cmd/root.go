package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "api-cli",
	Short: "A command-line tool for API development",
	Long:  `api-cli is a CLI tool for validating API schemas and linting code.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to api-cli!")
	},
}

func init() {
	RootCmd.AddCommand(schemaValidationCmd)
	RootCmd.AddCommand(lintCmd)
	RootCmd.AddCommand(apiValidationCmd)

	apiValidationCmd.AddCommand(validatorCmd)
	apiValidationCmd.AddCommand(nonameCmd)

	validatorCmd.PersistentFlags().StringP("file", "f", "", "the path to the file for api-validator (required)")
	nonameCmd.PersistentFlags().StringP("file", "f", "", "the path to the file for noname validation(required)")

	viper.BindPFlag("validatorFile", validatorCmd.PersistentFlags().Lookup("file"))
	viper.BindPFlag("nonameFile", nonameCmd.PersistentFlags().Lookup("file"))
}
