package route

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseRoute(t *testing.T) {
	dir := "./" // 指定要扫描的目录
	groups, err := scanDir(dir)
	if err != nil {
		fmt.Println("Error scanning directory:", err)
		return
	}

	var result []string
	printRoutes("/", groups, &result)

	fmt.Printf("| %-42s | %-6s | %-17s |\n", "Path", "Method", "Handler")
	fmt.Println(strings.Repeat("-", 68))
	for _, line := range result {
		fmt.Println(line)
	}
}
