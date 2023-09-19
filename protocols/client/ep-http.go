package client

import (
	"io/ioutil"
	"log"
	"net/http"
)

var client *http.Client

func EP_http_client() {

	client = http.DefaultClient

	resp, err := client.Get("http://127.0.0.1:8099")

	if err != nil {
		log.Printf("%v", err)
		return
	}

	defer resp.Body.Close()

	contentByte, err := ioutil.ReadAll(resp.Body)

	log.Println(string(contentByte))

}
