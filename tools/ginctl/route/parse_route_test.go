package route

import (
	"fmt"
	"testing"
)

func TestParseRoute(t *testing.T) {
	dir := "./" // 指定要扫描的目录
	groups, err := scanDir(dir)
	if err != nil {
		fmt.Println("Error scanning directory:", err)
		return
	}

	printRoutes(groups)
}
