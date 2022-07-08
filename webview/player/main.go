package main

import (
	"fmt"
	"log"
	"embed"
	"encoding/json"
	"path/filepath"
	"github.com/webview/webview"

	"local.package/api"
)

//go:embed static
var staticFS embed.FS

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Player")
	w.SetSize(800, 450, webview.HintNone)
	initBytes, err := staticFS.ReadFile(filepath.Join("static", "index.html"))
	if err != nil {
		log.Fatal(err)
	}
	w.Bind("setItemList", func(s string) {
		go func() {
			_ = api.SaveData(s)
		}
	})
	w.Bind("getItemList", func()(string) {
		jsonBytes, err = api.LoadData()
		if err != nil {
			fmt.Println(err)
			return "[]"
		} else {
			return string(jsonBytes)
		}
	})
	w.Bind("searchItemList", func(s string)(string) {
		data, err = api.Search(s)
		if err != nil {
			fmt.Println(err)
			return "[]"
		} else {
			jsonBytes, err_ := json.Marshal(data)
			if err_ != nil {
				fmt.Println(err)
				return "[]"
			} else {
				return string(jsonBytes)
			}
		}
	})
	w.Navigate("data:text/html," + url.PathEscape(string(initBytes)))
	w.Run()
}
