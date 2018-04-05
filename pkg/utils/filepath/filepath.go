package filepath

import (
	"strings"
)

// PKG returns file pkg
func PKG(path string) string {
	index := strings.Index(path, "/pkg")
	return path[index:]
}
