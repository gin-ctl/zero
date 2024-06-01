package route

import (
	"fmt"
	"github.com/gin-ctl/zero/package/helper"
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

func TestLines(t *testing.T) {
	routePath := "/Users/qinchaozheng/.g/zero/bootstrap/route.go"
	imports := fmt.Sprintf("\t\"github.com/gin-ctl/zero/app/http/%s/route\"", "demo")
	lines, er := helper.ReadLines(routePath)
	if er != nil {
		t.Error(er)
	}
	isExisted := helper.CheckLineIsExisted(lines, imports)
	t.Log(isExisted)
}
