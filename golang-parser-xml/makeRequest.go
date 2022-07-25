package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeRequest(a string, d []byte) {
	req, err := http.NewRequest("POST", a, bytes.NewBuffer(d))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
