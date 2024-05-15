package greenlight_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
)

func printResponseBody(resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Print(sb)
}

func printRequestBody(req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Print(sb)

}

// Unit tests for the API
func TestNewAccountCreationRequest(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Name":     "Rachel",
		"Email":    "rach@mail.ru",
		"Password": "password123",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/users", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestNewAccountCreationRequest2(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Name":     "Teodor",
		"Email":    "Breaking_Bad",
		"Password": "A1"})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/users", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestGettingAccountAuthenticationToken(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Email":    "Alibek@mail.ru",
		"Password": "password123",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/tokens/authentication", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestGettingAccountAuthenticationToken2(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Email":    "cccaaa@dot.net",
		"Password": "password123",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:4000/v1/tokens/authentication", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestActivatingAccount(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Token": "BPNT6BKDNMXT532QU7PTJJCV5Q",
	})
	responseBody := bytes.NewBuffer(postBody)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:4000/v1/users/activated", responseBody)
	if err != nil {
		log.Fatalf("An Error Occurred while creating request: %v", err)
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occurred while sending request: %v", err)
	}
	defer resp.Body.Close()

	// Print response body
	printResponseBody(resp)
}

func TestActivatingAccount2(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Token": " A VALID TOKEN", // ACTIVATION TOKEN
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.NewRequest(http.MethodPost, "http://localhost:4000/v1/users/activated", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	printRequestBody(resp)
}

//Integration tests for the API

type Movie struct {
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Runtime string   `json:"runtime"`
	Genres  []string `json:"genres"`
}

func TestInsertingMoviesIntoDatabase(t *testing.T) {
	moviePayload := Movie{
		Title:   "Interstellar",
		Year:    2010,
		Runtime: "210 min",
		Genres:  []string{"thriller"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "ZK633D2HGQGKZSCHQGWRZOLSYI"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestInsertingMoviesIntoDatabaseWithWrongYear(t *testing.T) {
	moviePayload := Movie{
		Title:   "Interstellar",
		Year:    2030,
		Runtime: "210 min",
		Genres:  []string{"fantasy"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "ZK633D2HGQGKZSCHQGWRZOLSYI"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestInsertingMoviesIntoDatabaseWithWrongRuntime(t *testing.T) {
	moviePayload := Movie{
		Title:   "Interstellar",
		Year:    2010,
		Runtime: "1030",
		Genres:  []string{"fantasy"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "ZK633D2HGQGKZSCHQGWRZOLSYI"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}

func TestMovieDeletionById(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:4000/v1/movies/5", nil)
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "ZK633D2HGQGKZSCHQGWRZOLSYI"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	printResponseBody(resp)
}
