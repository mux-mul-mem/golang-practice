package main

import (
	"os"
	"bytes"
	"strconv"
	"path/filepath"
)

func main() {
	basePath, _ := os.Getwd()
	sourcePath := filepath.Join(basePath, "..", "data", "base64.txt")
	resultPath := filepath.Join(basePath, "..", "data", "split01.txt")

	dataBytes, _ := os.ReadFile(sourcePath)
	var byteBuf bytes.Buffer
	var byteTmp byte
	num := 0
	for i, v := range dataBytes {
		if i == 0 {
			byteBuf.WriteByte(v)
		} else if i == len(dataBytes) - 1 {
			if v != byteTmp {
				byteBuf.WriteString(",")
				byteBuf.WriteString(strconv.Itoa(num))
				byteBuf.WriteString("\n")
				byteBuf.WriteByte(v)
				byteBuf.WriteString(",1")
			} else {
				byteBuf.WriteString(",")
				byteBuf.WriteString(strconv.Itoa(num + 1))
			}
		} else if v != byteTmp {
			byteBuf.WriteString(",")
			byteBuf.WriteString(strconv.Itoa(num))
			byteBuf.WriteString("\n")
			byteBuf.WriteByte(v)
			num = 0
		}
		byteTmp = v
		num += 1
	}

	os.WriteFile(resultPath, byteBuf.Bytes(), os.ModePerm)
}
