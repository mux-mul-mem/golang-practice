package main

import (
	"os"
	"path/filepath"
	"encoding/base64"
)

func main() {
	basePath, _ := os.Getwd()
	sourcePath := filepath.Join(basePath, "..", "data", "base64.txt")
	resultPath := filepath.Join(basePath, "..", "data", "favicon.ico")

	dataBytes, _ := os.ReadFile(sourcePath)

	decBytes, _ := base64.StdEncoding.DecodeString(string(dataBytes))

	os.WriteFile(resultPath, decBytes, os.ModePerm)
}
