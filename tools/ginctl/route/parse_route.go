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
	Method  string
	Path    string
	Handler string
}

// Group represents a group of routes
type Group struct {
	Path   string
	Routes []Route
}

// parseRoutes parses the routes from a file's AST
func parseRoutes(node *ast.File) []Group {
	var groups []Group
	var currentPath []string

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			if fun, ok := x.Fun.(*ast.SelectorExpr); ok {
				if fun.Sel.Name == GROUP {
					groupPath := getStringArg(x)
					currentPath = append(currentPath, groupPath)
				} else if isHTTPMethod(fun.Sel.Name) {
					method := fun.Sel.Name
					path := getStringArg(x)
					handler := getHandler(x)
					fullPath := buildFullPath(currentPath, path)
					groups = addRoute(groups, fullPath, method, handler)
				}
			}
			//case *ast.BlockStmt:
			//	if len(currentPath) > 0 {
			//		currentPath = currentPath[:len(currentPath)-1]
			//	}
		}
		return true
	})

	return groups
}

// addRoute adds a route to the appropriate group
func addRoute(groups []Group, fullPath, method, handler string) []Group {
	found := false
	for i, group := range groups {
		if group.Path == currentPathPrefix(fullPath) {
			groups[i].Routes = append(groups[i].Routes, Route{Method: method, Path: fullPath, Handler: handler})
			found = true
			break
		}
	}
	if !found {
		groups = append(groups, Group{Path: currentPathPrefix(fullPath), Routes: []Route{{Method: method, Path: fullPath, Handler: handler}}})
	}
	return groups
}

// buildFullPath constructs the full path based on the group stack
func buildFullPath(currentPath []string, path string) string {
	fullPath := "/" + strings.Join(append(currentPath, strings.TrimPrefix(path, "/")), "/")
	return strings.TrimSuffix(fullPath, "/")
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
	ch := make(chan []Group)

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

	var groups []Group
	for groupList := range ch {
		groups = append(groups, groupList...)
	}

	return groups, nil
}

// findMaxLengths finds the maximum length of each column for formatting purposes
func findMaxLengths(groups []Group) (int, int, int) {
	maxPathLength := len(Path)
	maxMethodLength := len(Method)
	maxHandlerLength := len(Handler)

	for _, group := range groups {
		for _, route := range group.Routes {
			pathLength := len(route.Path)
			methodLength := len(route.Method)
			handlerLength := len(route.Handler)

			if pathLength > maxPathLength {
				maxPathLength = pathLength
			}
			if methodLength > maxMethodLength {
				maxMethodLength = methodLength
			}
			if handlerLength > maxHandlerLength {
				maxHandlerLength = handlerLength
			}
		}
	}

	return maxPathLength, maxMethodLength, maxHandlerLength
}

// printRoutes prints the parsed routes in a tabular format
func printRoutes(groups []Group) {
	maxPathLength, maxMethodLength, maxHandlerLength := findMaxLengths(groups)

	fmt.Printf("| %-*s | %-*s | %-*s |\n", maxPathLength, Path, maxMethodLength, Method, maxHandlerLength, Handler)
	fmt.Printf("|-%s-|-%s-|-%s-|\n",
		strings.Repeat("-", maxPathLength),
		strings.Repeat("-", maxMethodLength),
		strings.Repeat("-", maxHandlerLength))

	for _, group := range groups {
		for _, route := range group.Routes {
			fmt.Printf("| %-*s | %-*s | %-*s |\n",
				maxPathLength, route.Path,
				maxMethodLength, route.Method,
				maxHandlerLength, route.Handler)
		}
	}
}

// currentPathPrefix returns the path prefix of the full path
func currentPathPrefix(fullPath string) string {
	parts := strings.Split(fullPath, "/")
	if len(parts) > 2 {
		return "/" + parts[1]
	}
	return fullPath
}
