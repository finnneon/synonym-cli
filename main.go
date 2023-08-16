package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Result struct {
	Meta struct {
		Syns [][]string `json:"syns"`
	} `json:"meta"`
}

func main() {
	apiKey := os.Getenv("THESAURUS_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set the THESAURUS_API_KEY environment variable to use. See README for instructions\n")
	}
	if len(os.Args) != 2 {
		log.Fatal("Usage: synonym-cli <word>\n")
	}
	word := os.Args[1]
	apiUrl := fmt.Sprintf("https://dictionaryapi.com/api/v3/references/thesaurus/json/%s?key=%s", word, apiKey)
	resp, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var resultStruct []Result
	json.Unmarshal(result, &resultStruct)
	for _, synonym := range resultStruct[0].Meta.Syns[0] {
		fmt.Println(synonym)
	}
}
