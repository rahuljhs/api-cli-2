package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nonameCmd = &cobra.Command{
	Use:   "noname",
	Short: "Validate API using noname",
	Long:  `This command validates API using noname`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Validating API schemas...")
	},
}
