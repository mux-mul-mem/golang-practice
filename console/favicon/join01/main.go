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
	sourcePath := filepath.Join(basePath, "..", "data", "split01.txt")
	resultPath := filepath.Join(basePath, "..", "data", "base64.txt")

	dataBytes, _ := os.ReadFile(sourcePath)
	var byteBuf bytes.Buffer
	for _, v := range regexp.MustCompile("[\n]").Split(string(dataBytes), -1) {
		if len(v) > 0 {
			x := regexp.MustCompile("[,]").Split(v, -1)
			if len(x) < 2 {
				continue
			}
			num, _ := strconv.Atoi(x[1])
			for i := 0; i < num; i++ {
				byteBuf.WriteString(x[0])
			}
		}
	}

	os.WriteFile(resultPath, byteBuf.Bytes(), os.ModePerm)
}
