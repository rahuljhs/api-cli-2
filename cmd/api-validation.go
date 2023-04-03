package cmd

import (
	"github.com/spf13/cobra"
)

var apiValidationCmd = &cobra.Command{
	Use:   "api-validation",
	Short: "Validate the API using api-validator and noname",
	Long:  `This command validates the API using api-validator and noname`,
}
