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
	Title         string   `json:"title"`
	Authors       []string `json:"authors"`
	Publisher     string   `json:"publisher"`
	PublishedDate string   `json:"publishedDate"`
	Categories    []string `json:"categories"`
	Description   string   `json:"description"`
}

type GoogleBooksResponse struct {
	Items []struct {
		VolumeInfo GoogleBooksVolumeInfo `json:"volumeInfo"`
	} `json:"items"`
}

type BookData struct {
	Title         string `json:"title"`
	Authors       string `json:"authors"`
	Publisher     string `json:"publisher"`
	PublishedYear string `json:"publishedYear"`
	Categories    string `json:"categories"`
	Description   string `json:"description"`
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

	authors := strings.Join(info.Authors, ", ")
	categories := strings.Join(info.Categories, ", ")
	publishedYear := ""
	if len(info.PublishedDate) >= 4 {
		publishedYear = info.PublishedDate[:4]
	}

	book := BookData{
		Title:         info.Title,
		Authors:       authors,
		Publisher:     info.Publisher,
		PublishedYear: publishedYear,
		Categories:    categories,
		Description:   info.Description,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
