package main

import (
	"io"
	"os"
	"fmt"
	"log"
	"bytes"
	"runtime"
	"os/exec"
	"net/http"
	"encoding/json"
	"path/filepath"

	"local.package/api"
)

func main() {
	basePath, err := os.Getwd()
	if err == nil {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				var path string
				if len(r.URL.Path[1:]) == 0 {
					path = filepath.Join(basePath, "static", "index.html")
				} else {
					path = filepath.Join(basePath, "static", r.URL.Path[1:])
				}
				http.ServeFile(w, r, path)
			} else {
				body := r.Body
				defer body.Close()
				buf := new(bytes.Buffer)
				io.Copy(buf, body)
				res, err := api.Handle(buf.Bytes())
				if err == nil {
					w.Header().Set("Content-Type", "application/json")
					jsonBytes, _ := json.Marshal(res)
					w.Write(jsonBytes)
				} else {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		})
		port := "8080"
		url := fmt.Sprintf("http://127.0.0.1:%s",port)
		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}
		if err == nil {
			http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}
