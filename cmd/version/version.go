package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Serve",
	Long:  `All software has versions. This is Serve's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Serve version %s\n", version)
	},
}
