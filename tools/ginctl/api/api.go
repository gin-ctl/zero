package api

import (
	"github.com/spf13/cobra"
	"sync"
)

var (
	apply string
	model string
	api   string
)

func GenerateApi() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model",
		Short: "make model",
		Long: `
		`,
		RunE: GenApi,
	}

	cmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
	cmd.Flags().StringVarP(&model, "model", "m", "", "Specify model name")
	cmd.Flags().StringVarP(&api, "api", "i", "", "Specify api name")

	return cmd
}

func GenApi(_ *cobra.Command, _ []string) (err error) {
	var wg sync.WaitGroup

	wg.Add(2)
	// logic
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

	}(&wg)

	// types
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

	}(&wg)
	wg.Wait()

	return
}
