package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var schemaValidationCmd = &cobra.Command{
	Use:   "schema-validation",
	Short: "Validate API schemas",
	Long:  `The schema-validation command validates API schemas according to a specified schema format.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Validating API schemas...")
	},
}
