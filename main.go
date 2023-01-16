package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func translate_google(sourceLang string, targetLang string, queryText string) {
	requestUrl := "https://translate.google.com/translate_a/single?client=at&dt=t&dt=ld&dt=qca&dt=rm&dt=bd&dj=1&ie=UTF-8&oe=UTF-8&inputm=2&otf=2&iid=1dd3b944-fa62-4b55-b330-74909a99969e"

	formData := url.Values{
		"sl": {sourceLang},
		"tl": {targetLang},
		"q":  {queryText},
	}

	resp, err := http.PostForm(requestUrl, formData)
	if err != nil {
		log.Fatalf("An Error occured %v", err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println(result["sentences"])
}

func cmd_line_arguments() {
	sourceLang := os.Args[1]
	targetLang := os.Args[2]
	query := os.Args[3]

	translate_google(sourceLang, targetLang, query)
}

func run_with_input() {
	var sourceLang, targetLang, query string

	fmt.Print("Enter source language ")
	fmt.Scanln(&sourceLang)
	fmt.Print("Enter target language ")
	fmt.Scanln(&targetLang)
	fmt.Print("Enter query string ")
	fmt.Scanln(&query)

	translate_google(sourceLang, targetLang, query)
}

func test() {
	f, err := os.Open("source.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		query := scanner.Text()
		fmt.Printf("translating: %s\n", query)
		translate_google("en", "fr", query)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//run_with_input()
	cmd_line_arguments()

	// Uncomment the line below and comment out the above line to test the code with the source and target text files
	//test()
}
