package mget

import (
	"path/filepath"
	"strings"
)

// parseDirAndFileName returns dir path and file name from given path
func parseDirAndFileName(path string) (dir, file string) {
	lastSlashIdx := strings.LastIndex(path, "/")
	dir = path[:lastSlashIdx+1]
	if len(dir) == len(path) { // path is a dir
		return dir, ""
	}

	return dir, path[(len(dir) + 1):]
}

// parseFileName returns file name from given url
func parseFileName(url string) string {
	return filepath.Base(url)
}
