package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://docs.aws.amazon.com/sitemap_index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	resp.Body.Close()

	fmt.Println(stringBody)
}
