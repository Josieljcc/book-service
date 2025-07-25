package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type BookResponse struct {
	Message string `json:"message"`
}

type GoogleBooksVolumeInfo struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
}

type GoogleBooksResponse struct {
	Items []struct {
		VolumeInfo GoogleBooksVolumeInfo `json:"volumeInfo"`
	} `json:"items"`
}

type BookData struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	// Espera-se que a URL seja /book/{isbn}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[1] != "book" || parts[2] == "" {
		http.Error(w, "ISBN not provided", http.StatusBadRequest)
		return
	}
	isbn := parts[2]

	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")
	if apiKey == "" {
		http.Error(w, "API key not configured", http.StatusInternalServerError)
		return
	}
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%s&key=%s", isbn, apiKey)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		http.Error(w, "Error querying Google Books API", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading API response", http.StatusInternalServerError)
		return
	}

	var gbResp GoogleBooksResponse
	err = json.Unmarshal(body, &gbResp)
	if err != nil || len(gbResp.Items) == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	info := gbResp.Items[0].VolumeInfo
	book := BookData(info)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
