package main

import (
	"os"
	"path/filepath"
	"encoding/base64"
)

func main() {
	basePath, _ := os.Getwd()
	sourcePath := filepath.Join(basePath, "..", "data", "favicon.ico")
	resultPath := filepath.Join(basePath, "..", "data", "base64.txt")

	dataBytes, _ := os.ReadFile(sourcePath)

	encString := base64.StdEncoding.EncodeToString(dataBytes)

	os.WriteFile(resultPath, []byte(encString), os.ModePerm)
}
