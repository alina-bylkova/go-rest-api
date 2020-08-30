package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getAll() {
	resp, err := http.Get("http://localhost:8080/numbers")
	if err != nil {
		fmt.Println(err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func post(num int) {
	reqBody, err := json.Marshal(map[string]int{
		"num": num,
	})

	if err != nil {
		fmt.Println(err)
	}
	resp, err := http.Post("http://localhost:8080/numbers",
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func main() {
	for i := 0; i < 50; i++ {
		go post(i)
		go getAll()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("ended")
}
