package mycmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RunServe the main event loop for the service
func RunMyCommand() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		/* Using flags
		myflag, err := cmd.Flags().GetString("myflag")
		if err != nil {
			log.Warn().Msg("Missing myflag")
			os.Exit(1)
		}
		app.Env.MyFlag = myflag
		*/

		fmt.Println("Hello World!")

		os.Exit(0)
	}
}
