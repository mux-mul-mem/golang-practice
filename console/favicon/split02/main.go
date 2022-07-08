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
	resultPath := filepath.Join(basePath, "..", "data", "split02.txt")

	dataBytes, _ := os.ReadFile(sourcePath)
	var byteBuf bytes.Buffer
	var byteTmp byte
	num := 0
	for i, v := range dataBytes {
		if i == 0 {
			byteBuf.WriteByte(v)
		} else if i == len(dataBytes) - 1 {
			if v != byteTmp {
				byteBuf.WriteString(strconv.FormatUint(uint64(num), 16))
				byteBuf.WriteString(",")
				byteBuf.WriteByte(v)
				byteBuf.WriteString("1")
			} else {
				byteBuf.WriteString(strconv.FormatUint(uint64(num + 1), 16))
			}
		} else if v != byteTmp {
			byteBuf.WriteString(strconv.FormatUint(uint64(num), 16))
			byteBuf.WriteString(",")
			byteBuf.WriteByte(v)
			num = 0
		}
		byteTmp = v
		num += 1
	}

	os.WriteFile(resultPath, byteBuf.Bytes(), os.ModePerm)
}
