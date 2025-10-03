package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"google.golang.org/genai"
	"google.golang.org/api/option"
)

// Message represents the structure of the incoming chat request.
type Message struct {
	Prompt string `json:"prompt"`
}

// ChatResponse represents the structure of the outgoing chat response.
type ChatResponse struct {
	Response string `json:"response"`
}

func main() {
	http.HandleFunc("/chat", chatHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Message
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Prompt == "" {
		http.Error(w, "Prompt cannot be empty", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		http.Error(w, "GEMINI_API_KEY environment variable not set", http.StatusInternalServerError)
		return
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		http.Error(w, "Failed to create Gemini client", http.StatusInternalServerError)
		log.Printf("Failed to create Gemini client: %v", err)
		return
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	resp, err := model.GenerateContent(ctx, genai.Text(req.Prompt))
	if err != nil {
		http.Error(w, "Failed to generate content from Gemini API", http.StatusInternalServerError)
		log.Printf("Failed to generate content: %v", err)
		return
	}

	var geminiResponse string
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		geminiResponse = string(resp.Candidates[0].Content.Parts[0].(genai.Text))
	} else {
		geminiResponse = "No response from AI."
	}

	responseJSON := ChatResponse{Response: geminiResponse}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseJSON)
}
