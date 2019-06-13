package opt

import "strings"

func OptPath(paths []string) []string {
	dirlist := make([]string, 0)
	for _, p := range paths {
		if !contains(dirlist, p) {
			dirlist = append(dirlist, p)
		}
	}

	return dirlist
}

func contains(s []string, e string) bool {
	if len(s) == 0 {
		return false
	}
	for _, v := range s {
		if strings.HasPrefix(v, e) || strings.HasPrefix(e, v) {
			return true
		}
	}
	return false
}
