package target

import (
	"os"
	"path/filepath"
)

// Get root下のin形式のファイルのパスが入ったsliceとerrを返す
func Get(root, in string) ([]string, error) {
	targets := []string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+in {
			targets = append(targets, path)
		}

		return nil
	})

	if err != nil {
		return targets, err
	}

	return targets, nil
}
