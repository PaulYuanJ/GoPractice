package main

import (
	"fmt"
	progressbar "github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
)

//var env string

func init() {
	fmt.Println("I am init")
	//flag.StringVar(&env, "env", "release", "This is a help")
}

func main() {
	fmt.Println("aaaaaaaaaa")
	url := "https://oss.xcloud.lenovo.com:8082/awp-public/hololib-linux-83c5f83e56c2cce1.zip"
	hash := "83c5f83e56c2cce1"

	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	f, _ := os.OpenFile(fmt.Sprintf("hololib-%s.zip", hash), os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	fmt.Println(resp.ContentLength)
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
}
