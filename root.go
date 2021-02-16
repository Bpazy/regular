package regular

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// buildVer represents 'regular' build version
	buildVer string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "regular",
		Short: "TODO",
		Long: `TODO
`,
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "版本号",
		Long:  `查看 regular 的版本号`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(buildVer)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
