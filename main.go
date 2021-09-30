package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BookRange struct {
	ID    int64 `json:"id"`
	Start int   `json:"start"`
	End   int   `json:"end"`
}

type Book struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	BookRange
}

const URL = "http://localhost:3000/add"

func main() {

	sample := Book{
		ID:      1,
		Title:   "タイトル",
		Content: "コンテンツ",
		BookRange: BookRange{
			ID:    5,
			Start: 4,
			End:   6,
		},
	}

	// encode json
	sample_json, _ := json.Marshal(sample)
	fmt.Printf("[+] %s\n", string(sample_json))

	// send json
	//// ポイント2, 3
	res, err := http.Post(URL, "application/json", bytes.NewBuffer(sample_json))
	defer res.Body.Close()

	if err != nil {
		fmt.Println("[!] " + err.Error())
	} else {
		fmt.Println("[*] " + res.Status)
	}
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("[body] " + string(body))

}
