package api

import (
	"fmt"
	"github.com/gin-ctl/zero/package/console"
	"github.com/spf13/cobra"
	"strings"
	"sync"
)

var OptMap = []string{
	"Index",
	"Show",
	"Create",
	"Update",
	"Destroy",
}

var apiCmd = &cobra.Command{
	Use:   "opt",
	Short: "make operation",
	Long:  ``,
	RunE:  GenOperation,
}

func GenOperation(_ *cobra.Command, _ []string) (err error) {

	if operation != "" && curd {
		console.Exit("Custom operations cannot be specified at the same time as CURD operations.")
		return
	}

	if apply == "" {
		console.Error("invalid apply name.")
		return
	}

	if model == "" {
		console.Error("invalid model name.")
		return
	}

	files := map[StubCode]string{
		FromStubImport: "logic",
		FromStubTypes:  "types",
	}
	var wg sync.WaitGroup
	errChan := make(chan error, 10)
	body := &Body{
		LowerModel: strings.ToLower(model),
		Apply:      strings.ToLower(apply),
	}

	wg.Add(len(files))
	for code, file := range files {
		go func(wg *sync.WaitGroup, code StubCode, file string) {
			defer wg.Done()
			filePath := fmt.Sprintf("app/http/%s/logic/%s/%s.go", body.Apply, body.LowerModel, file)
			err = GenLogic(filePath, code, ToLogic, body)
			if err != nil {
				errChan <- err
			}
		}(&wg, code, file)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for errs := range errChan {
		if err != nil {
			console.Error(errs.Error())
			return
		}
	}

	// make chan
	//opt := make(chan )
	if curd {
		// CURD
		for _, s := range OptMap {
			DoGenOperation(s)
		}
	} else {
		// Custom
		if operation == "" {
			console.Error("invalid operation name.")
			return
		}
		DoGenOperation(operation)
	}

	console.Success("Done.")

	return
}

func DoGenOperation(opt string) {

}
