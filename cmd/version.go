package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version      = "1.0.0"
	completion   bool
	contactName  = "API Team"
	contactEmail = "api-team@example.com"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of api-cli",
	Long:  `All software has versions. This is api-cli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("api-cli version %s\n%s\n", Version, contactDetails())
	},
}

func contactDetails() string {
	return fmt.Sprintf("%s (%s)", contactName, contactEmail)
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.PersistentFlags().BoolVarP(&completion, "generate-bash-completion", "g", false, "Generate bash completion")

	RootCmd.SetVersionTemplate("api-cli version {{.Version}}\n")
	RootCmd.CompletionOptions.DisableDefaultCmd = true
	RootCmd.CompletionOptions.DisableNoDescFlag = true
}
