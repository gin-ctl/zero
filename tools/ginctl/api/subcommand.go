package api

import (
	"fmt"
	"github.com/gin-ctl/zero/package/console"
	"github.com/spf13/cobra"
	"strings"
	"sync"
)

type Files struct {
	Code StubCode
	Name string
}

type Opts struct {
	Name string
	Desc string
}

var apiCmd = &cobra.Command{
	Use:   "opt",
	Short: "make operation",
	Long: `Example: api opt -a test -m user -c true CURD operation to create a resource. 
Example: api opt -a test -m user -o ping to create a single operation method.`,
	RunE: GenOperation,
}

func init() {
	apiCmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
	apiCmd.Flags().StringVarP(&model, "model", "m", "", "Specify model name")
	apiCmd.Flags().StringVarP(&operation, "operation", "o", "", "Specify operation name")
	apiCmd.Flags().StringVarP(&desc, "desc", "d", "", "Specify operation description")
	apiCmd.Flags().BoolVarP(&curd, "curd", "c", false, "Specifies whether you need to generate add, delete, update and get operations for the module")
}

func GenOperation(_ *cobra.Command, args []string) (err error) {
	Cmd.Run(Cmd, args)

	if operation != "" && curd {
		console.Error("Custom operations cannot be specified at the same time as CURD operations.")
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

	files := []Files{
		{
			Code: FromStubImport,
			Name: "logic",
		},
		{
			Code: FromStubTypes,
			Name: "types",
		},
	}
	var wg sync.WaitGroup
	errChan := make(chan error, 10)
	body := &Body{
		LowerModel: strings.ToLower(model),
		Apply:      strings.ToLower(apply),
	}

	wg.Add(len(files))
	for _, info := range files {
		go func(wg *sync.WaitGroup, code StubCode, file string) {
			defer wg.Done()
			filePath := fmt.Sprintf("app/http/%s/logic/%s/%s.go", body.Apply, body.LowerModel, file)
			err = GenLogic(filePath, code, ToLogic, body)
			if err != nil {
				errChan <- err
			}
		}(&wg, info.Code, info.Name)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for errs := range errChan {
		if errs != nil {
			console.Error(errs.Error())
			return
		}
	}

	// make operation.
	errs := make(chan error, 20)
	files = []Files{
		{
			Code: FromStubLogicFunc,
			Name: "logic",
		},
		{
			Code: FromStubTypeFunc,
			Name: "types",
		},
	}
	if curd {
		// CURD
		OptMap := []Opts{
			{
				Name: "Index",
				Desc: "Get page list",
			},
			{
				Name: "Show",
				Desc: "Get info",
			},
			{
				Name: "Create",
				Desc: "Save one source",
			},
			{
				Name: "Update",
				Desc: "Modifying a resource",
			},
			{
				Name: "Destroy",
				Desc: "Delete a resource",
			},
		}
		for _, opt := range OptMap {
			for _, info := range files {
				filePath := fmt.Sprintf("app/http/%s/logic/%s/%s.go", body.Apply, body.LowerModel, info.Name)
				DoGenOperation(filePath, opt.Name, opt.Desc, info.Code, errs)
			}
		}
	} else {
		// Custom
		if operation == "" {
			console.Error("invalid operation name.")
			return
		}
		for _, info := range files {
			filePath := fmt.Sprintf("app/http/%s/logic/%s/%s.go", body.Apply, body.LowerModel, info.Name)
			if desc == "" {
				desc = operation
			}
			DoGenOperation(filePath, operation, desc, info.Code, errs)
		}
	}

	close(errs)
	for err = range errs {
		if err != nil {
			console.Error(err.Error())
			return
		}
	}

	console.Success("Done.")

	return
}
