package walk

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Walker interface {
	Find(string, string) ([]string, error)
}

type walker struct{}

func NewWalker() Walker {
	return &walker{}
}

func (w *walker) Find(path string, ext string) ([]string, error) {
	numFile := 0
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() &&
			strings.TrimLeft(filepath.Ext(info.Name()), ".") == ext {
			numFile++
			fmt.Println(info.Name())
		}
		return nil
	})

	fmt.Println(numFile)

	paths := make([]string, numFile)
	i := 0
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() &&
			strings.TrimLeft(filepath.Ext(info.Name()), ".") == ext {
			paths[i] = path
			i++
		}
		return nil
	})

	return paths, nil
}
