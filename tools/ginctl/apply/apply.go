package apply

import (
	"github.com/gin-ctl/zero/package/console"
	"github.com/spf13/cobra"
)

var (
	name   string
	module string
	flags  string
)

var Cmd = &cobra.Command{
	Use:   "app",
	Short: "make apply",
	Long:  `Generate an http api application.`,
	RunE:  GenHttpApply,
}

const (
	Api       = "api"
	Grpc      = "grpc"
	Websocket = "websocket"
)

// By add subcommand order.
var subCmdMap = map[string]int{
	Api: 0,
	//Grpc:      1,
	//Websocket: 2,
}

func init() {
	Cmd.Flags().StringVarP(&name, "name", "n", "", "Specify apply name")
	Cmd.Flags().StringVarP(&module, "module", "m", "api", "Specify module name")
	api.Flags().StringVarP(&flags, "flags", "f", "api", "Specify module name")
	Cmd.AddCommand(api)
}

func GenHttpApply(cmd *cobra.Command, _ []string) (err error) {

	if name == "" {
		console.Error("apply name invalid.")
		return
	}

	if module == "" {
		console.Error("module name invalid.")
		return
	}

	index, exists := subCmdMap[module]
	if !exists {
		console.Error("module name not found.")
		return
	}

	subCmd := cmd.Commands()[index]
	subCmd.SetArgs([]string{module, "-f", name})
	//switch module {
	//case Api:
	//
	//case Grpc:
	//case Websocket:
	//default:
	//	console.Error("unsupported module.")
	//	return nil
	//}

	if subCmd != nil && subCmd != cmd {
		err = subCmd.Execute()
		if err != nil {
			console.Error(err.Error())
			return
		}
	}

	console.Success("Done.")
	return
}
