package interview150

import (
	"fmt"
	"strings"
)

func SimplifyPath() {
	fmt.Println(simplifyPath("/.../"))
}

func simplifyPath(path string) string {
	var result []string
	list := strings.Split(path, "/")
	for _, v := range list {
		switch v {
		case "", ".":
			continue
		case "..":
			if len(result) == 0 {
				continue
			}
			result = result[:len(result)-1]
		default:
			fmt.Println(v)
			result = append(result, "/"+v)
		}
	}
	if len(result) == 0 {
		return "/"
	}
	return strings.Join(result, "")
}
