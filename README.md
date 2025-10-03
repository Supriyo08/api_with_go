# Go Gemini Chat API

This is a simple and lightweight backend API built with Go that allows you to chat with the Gemini AI model.  
It provides a single REST endpoint to send a text prompt and receive a generated response from the Gemini API.

## Features
- **Go Standard Library**: Uses Go's built-in `net/http` (no external frameworks).
- **Gemini API Integration**: Communicates with Google's Gemini AI via the official Go client.
- **Environment Variables**: API key is stored securely in environment variables.
- **JSON-based API**: Standard request/response structure.

---

## Prerequisites
- Go 1.18 or higher → [Install Go](https://go.dev/dl/)
- Gemini API Key → Get it from [Google AI Studio](https://aistudio.google.com/)

---

## Setup

### 1. Clone Project
```bash
git clone https://github.com/yourusername/go-gemini-chat-api.git
cd go-gemini-chat-api
```

### 2. Initialize Go Module
```bash
go mod init go-gemini-chat-api
```

### 3. Install Dependencies
```bash
go get google.golang.org/genai
```

### 4. Set API Key
Linux / macOS:
```bash
export GEMINI_API_KEY="your_api_key_here"
```

Windows (PowerShell):
```powershell
$env:GEMINI_API_KEY="your_api_key_here"
```

---

## Running the Server
```bash
go run main.go
```
Server will start at:  
👉 [http://localhost:8080](http://localhost:8080)

---

## API Endpoint

### `POST /chat`
**Request:**
```json
{
  "prompt": "What is the capital of France?"
}
```

**Response:**
```json
{
  "response": "The capital of France is Paris."
}
```

---

## Test with curl
```bash
curl -X POST http://localhost:8080/chat -H "Content-Type: application/json" -d '{"prompt": "Tell me a fun fact"}'
```

---

## Error Responses
- `400` → Invalid request body or empty prompt  
- `405` → Method not allowed  
- `500` → Missing API key or Gemini API failure  

---

## License
MIT License
