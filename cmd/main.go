package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

func main() {
	downloads := xdg.UserDirs.Download
	ext := findExtensions(downloads)
	fmt.Println(ext)
	result, erro := FindFilesByExtension(downloads, "")
	if erro != nil {
		fmt.Println("Erro", erro)
	}
	for _, results := range result {
		fmt.Println(results)
	}
}

func FindFilesByExtension(dir, ext string) ([]string, error) {
	var files []string

	err := filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && filepath.Ext(path) == ext {
				files = append(files, path)
			}

			return nil
		},
	)

	return files, err
}

func findExtensions(path string) string {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		parts := strings.Split(file.Name(), ".")
		fmt.Println(parts)
	}
	return ""
}
