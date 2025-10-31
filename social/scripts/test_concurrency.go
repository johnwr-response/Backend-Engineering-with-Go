package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type UpdatePostPayload struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func updatePost(postID int, p UpdatePostPayload, wg *sync.WaitGroup) {
	defer wg.Done()

	// Construct the URL for the update endpoint
	url := fmt.Sprintf("http://localhost:8080/v1/posts/%d", postID)

	// Create the JSON payload
	b, _ := json.Marshal(p)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers as needed:
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println("Error closing body:", err)
		}
	}(resp.Body)

	fmt.Println("Update response status:", resp.Status)
}
func main() {
	var wg sync.WaitGroup

	// Assuming the post ID to update is 39
	postID := 39

	// Simulate user A and user B updating the same post concurrently
	wg.Add(2)
	content := "NEW CONTENT FROM USER B"
	title := "NEW TITLE FROM USER A"

	go updatePost(postID, UpdatePostPayload{Title: &title}, &wg)
	go updatePost(postID, UpdatePostPayload{Content: &content}, &wg)
	wg.Wait()

}
