package apply

import (
	"github.com/gin-ctl/zero/package/console"
	"github.com/spf13/cobra"
)

var api = &cobra.Command{
	Use:   "api",
	Short: "make api",
	Long:  `Generate an http api application.`,
	RunE:  GenModuleApply,
}

func GenModuleApply(cmd *cobra.Command, args []string) (err error) {
	console.Success("Yes !")
	return
}
