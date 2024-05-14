package main

import "github.com/spf13/cobra"

func main() {

	var rootCmd = &cobra.Command{
		Use:   "zero",
		Short: "gc",
		Long:  `可以用 -h 查看更多命令`,
	}

	rootCmd.AddCommand(
	//cmd.MakeAPI(),
	//cmd.MakeModel(),
	)

	// 执行命令
	cobra.CheckErr(rootCmd.Execute())

}
