package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, errGet := http.Get("http://localhost:8795/time")
	if errGet != nil {
		fmt.Println(errGet)
		return
	}
	defer resp.Body.Close()
	body, errResp := ioutil.ReadAll(resp.Body)
	if errResp != nil {
		fmt.Println(errResp)
	}
	fmt.Println(string(body))
}

