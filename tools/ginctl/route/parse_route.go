package route

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	GROUP   = "Group"
	Path    = "Path"
	Method  = "Method"
	Handler = "Handler"
)

// Route represents a single route with method, path, and handler
type Route struct {
	Path    string
	Method  string
	Handler string
}

// Group represents a group of routes
type Group struct {
	Path   string
	Groups []Group
	Routes []Route
}

// parseRoutes parses the routes from a file's AST
func parseRoutes(node *ast.File) []Group {
	var groups []Group

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			if fun, ok := x.Fun.(*ast.SelectorExpr); ok {
				group := Group{}
				switch {
				case fun.Sel.Name == GROUP:
					group.Path = getStringArg(x)
				case isHTTPMethod(fun.Sel.Name):
					method := fun.Sel.Name
					path := getStringArg(x)
					handler := getHandler(x)
					group.Routes = append(group.Routes, Route{
						Path:    path,
						Method:  method,
						Handler: handler,
					})
				}
				groups = append(groups, group)
			}
		}
		return true
	})

	return groups
}

// getStringArg gets the first string argument from a call expression
func getStringArg(callExpr *ast.CallExpr) string {
	for _, arg := range callExpr.Args {
		if lit, ok := arg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			return strings.Trim(lit.Value, "\"")
		}
	}
	return ""
}

// getHandler gets the handler function from a call expression
func getHandler(callExpr *ast.CallExpr) string {
	if len(callExpr.Args) > 0 {
		if fun, ok := callExpr.Args[len(callExpr.Args)-1].(*ast.SelectorExpr); ok {
			if x, okk := fun.X.(*ast.Ident); okk {
				return fmt.Sprintf("%s.%s", x.Name, fun.Sel.Name)
			}
			return fun.Sel.Name
		}
	}
	return ""
}

// isHTTPMethod checks if a method is an HTTP method
func isHTTPMethod(method string) bool {
	switch method {
	case "GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD":
		return true
	default:
		return false
	}
}

// parseFile parses a Go source file and sends the parsed groups to a channel
func parseFile(filename string, wg *sync.WaitGroup, ch chan<- []Group) {
	defer wg.Done()

	fSet := token.NewFileSet()
	node, err := parser.ParseFile(fSet, filename, nil, parser.ParseComments)
	if err != nil {
		return
	}

	groups := parseRoutes(node)
	ch <- groups
}

// scanDir scans a directory for Go files and parses them concurrently
func scanDir(dir string) ([]Group, error) {
	var wg sync.WaitGroup
	ch := make(chan []Group, 10)

	var rootGroups []Group

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			wg.Add(1)
			go parseFile(path, &wg, ch)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for group := range ch {
		rootGroups = append(rootGroups, group...)
	}

	return rootGroups, nil
}

// printRoutes
func printRoutes(prefix string, groups []Group, result *[]string) {
	for _, group := range groups {
		newPrefix := prefix + group.Path
		for _, route := range group.Routes {
			path := newPrefix
			if route.Path != "" {
				path += "/" + route.Path
			}
			*result = append(*result, fmt.Sprintf("| %-42s | %-6s | %-17s |", path, route.Method, route.Handler))
		}
		printRoutes(newPrefix+"/", group.Groups, result)
	}
}
