package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint API code",
	Long:  `The lint command checks the quality of the code according to best practices and standards.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Linting API code...")
	},
}
