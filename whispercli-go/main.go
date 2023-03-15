package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: whisper-cli [path to audio file]")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create HTTP client
	client := &http.Client{}

	// Build API endpoint URL
	apiEndpoint := "https://whisper.ai/api/transcribe"
	formData := url.Values{}
	formData.Add("apikey", "your_api_key_here")
	formData.Add("model", "default")
	formData.Add("format", "txt")
	formData.Add("content-type", "audio/mpeg")

	// Build HTTP request
	req, err := http.NewRequest("POST", apiEndpoint, file)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "audio/mpeg")
	req.Header.Set("User-Agent", "Whisper-CLI/1.0")
	req.PostForm = formData

	// Send HTTP request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	// Check response status code
	if res.StatusCode != http.StatusOK {
		fmt.Println("API error:", res.Status)
		return
	}

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Output transcribed notes
	notes := strings.TrimSpace(string(body))
	fmt.Println(notes)
}
