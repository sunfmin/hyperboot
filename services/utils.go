package services

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"path/filepath"
)

func createFileWithContentSkipExists(relativePath string, fullPath string, content string, withoutExistOutput bool) (err error) {

	if _, err = os.Stat(fullPath); err == nil {
		if !withoutExistOutput {
			fmt.Printf("File exists %s, skipping.\n", relativePath)
		}
		return
	}

	err = os.MkdirAll(filepath.Dir(fullPath), 0755)
	if err != nil {
		fmt.Printf("Mkdir %s error: %s.\n", fullPath, err)
		return
	}

	var file *os.File
	file, err = os.Create(fullPath)
	if err != nil {
		fmt.Printf("Creating %s error: %s\n", relativePath, err)
		return
	}
	defer file.Close()
	var fsource []byte
	fsource, err = format.Source([]byte(content))
	if err != nil {
		fmt.Printf("Format source code error: %s \n", err)
		return
	}
	_, err = io.Copy(file, bytes.NewBuffer(fsource))
	fmt.Printf("create %s\n", relativePath)
	return
}
