package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)
var (
	supabaseUrl string
	supabaseKey string
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: No .env file found, using system environment variables.")
	}
	supabaseUrl = os.Getenv("SUPABASE_URL")
	supabaseKey = os.Getenv("SUPABASE_KEY")
}

func FetchFromSupabase() ([]byte, error) {

	if supabaseUrl == "" || supabaseKey == "" {
		return nil, fmt.Errorf("SUPABASE_URL or SUPABASE_KEY environment variable not set")
	}

	url := fmt.Sprintf("%s/rest/v1/%s?select=*", supabaseUrl, "profile")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Authorization", "Bearer "+supabaseKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// func CompareResponse 
// compares two texts passing through 7 layers of python microservice 
// return a score 
// indicating if this is a fraud or not
func CompareResponse(question,text1, text2 string) (int, error) {
	if text1 == "" || text2 == "" {
		return 0, fmt.Errorf("text1 or text2 is empty")
	}

	// Prepare the command to call the Python microservice
	cmd := exec.Command("python", "microservice/pipeline.py", question, text1, text2)

	// Capture the output
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("failed to execute python microservice: %v", err)
	}

	// Parse the output as an integer score
	var score int
	_, err = fmt.Sscanf(string(output), "%d", &score)
	if err != nil {
		return 0, fmt.Errorf("failed to parse score from output: %v", err)
	}

	return score, nil
}