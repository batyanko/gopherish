package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/batyanko/gopherish/pkg/translator"
)

type wordRequest struct {
	EnglishWord string `json:"english-word"`
}

type wordResponse struct {
	GopherWord string `json:"gopher-word"`
}

type sentenceRequest struct {
	EnglishSentence string `json:"english-sentence"`
}

type sentenceResponse struct {
	GopherSentence string `json:"gopher-sentence"`
}

type history struct {
	History []map[string]string `json:"history"`
}

// translateWord handles single word translation.
func (hist *history) translateWord(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, fmt.Sprintf("Method %s not supported", r.Method), http.StatusBadRequest)
		return
	}

	var request wordRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := wordResponse{
		GopherWord: translator.TranslateWord(request.EnglishWord),
	}

	// Update translation history.
	hist.History = append(hist.History, map[string]string{request.EnglishWord: resp.GopherWord})

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(js); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// translateSentence handles sentence translation.
func (hist *history) translateSentence(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, fmt.Sprintf("Method %s not supported", r.Method), http.StatusBadRequest)
		return
	}

	var request sentenceRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sentence, err := translator.TranslateSentence(request.EnglishSentence)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := sentenceResponse{
		GopherSentence: sentence,
	}

	// Update translation history.
	hist.History = append(hist.History, map[string]string{request.EnglishSentence: resp.GopherSentence})

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(js); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hist *history) listHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Method %s not supported", r.Method), http.StatusBadRequest)
		return
	}

	// Time to sort history of translations...
	sort.Slice(hist.History, func(i, j int) bool {
		var engI, _ string
		for engI = range hist.History[i] {
		}
		var engJ, _ string
		for engJ = range hist.History[j] {
		}

		return strings.ToLower(engI[:1]) < strings.ToLower(engJ[:1])
	})

	js, err := json.Marshal(hist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(js); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	arguments := os.Args[1:]
	help := flag.Bool("h", false, "show this help screen")
	flag.Parse()

	if *help || len(arguments) == 0 {
		fmt.Printf("Use this tool to start the server that translates English into Gopherish.\n")
		fmt.Printf("Refer to https://github.com/batyanko/gopherish for instructions on using the server.\n\n")
		fmt.Printf("Arguments:\n")
		fmt.Printf("The only accepted argument is port number for the HTTP server.\n")

		return
	}

	if len(arguments) > 1 {
		fmt.Printf("Too many arguments.\n")
		return
	}

	port := arguments[0]
	// Test if provided arg is valid port number in the range of an unsigned uint16
	_, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		fmt.Printf("Cannot use %s as listen port.\n", port)
		return
	}

	hist := history{History: []map[string]string{}}

	http.HandleFunc("/word", hist.translateWord)
	http.HandleFunc("/sentence", hist.translateSentence)
	http.HandleFunc("/history", hist.listHistory)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
