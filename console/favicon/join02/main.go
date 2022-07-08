package main

import (
	"os"
	"bytes"
	"regexp"
	"strconv"
	"path/filepath"
)

func main() {
	basePath, _ := os.Getwd()
	sourcePath := filepath.Join(basePath, "..", "data", "split02.txt")
	resultPath := filepath.Join(basePath, "..", "data", "base64.txt")

	dataBytes, _ := os.ReadFile(sourcePath)
	var byteBuf bytes.Buffer
	for _, v := range regexp.MustCompile("[,]").Split(string(dataBytes), -1) {
		if len(v) > 1 {
			num, _ := strconv.ParseUint(v[1:], 16, 0)
			for i := 0; i < int(num); i++ {
				byteBuf.WriteString(v[:1])
			}
		}
	}

	os.WriteFile(resultPath, byteBuf.Bytes(), os.ModePerm)
}
