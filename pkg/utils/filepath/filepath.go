package filepath

import (
	"strings"
)

// PKG returns file pkg
func PKG(path string) string {
	var (
		pattern = "/pkg"
		length  = len(pattern)
	)

	index := strings.Index(path, pattern)

	return path[index+length:]
}
