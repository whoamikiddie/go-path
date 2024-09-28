package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// Prompt the user for a URL
	fmt.Print("Enter the URL: ")
	reader := bufio.NewReader(os.Stdin)
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	// Prompt the user for the number of requests
	fmt.Print("Enter the number of requests: ")
	var count int
	if _, err := fmt.Scanf("%d", &count); err != nil {
		fmt.Println("Error reading the number of requests:", err)
		return
	}

	// Prompt the user for timeout
	fmt.Print("Enter the timeout in seconds: ")
	var timeout int
	if _, err := fmt.Scanf("%d", &timeout); err != nil {
		fmt.Println("Error reading the timeout:", err)
		return
	}

	// Create a custom transport
	transport := &http.Transport{
		MaxIdleConns:    100,
		IdleConnTimeout: 30 * time.Second,
	}

	// Create a shared HTTP client with user-defined timeout
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(timeout) * time.Second,
	}

	// List of user agents
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		// Add more user agents as needed
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Prompt for rate limiting
	fmt.Print("Enter the delay between requests in seconds: ")
	var delayBetweenRequests int
	if _, err := fmt.Scanf("%d", &delayBetweenRequests); err != nil {
		fmt.Println("Error reading the delay:", err)
		return
	}

	// Open a log file for errors
	logFile, err := os.OpenFile("error_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// Send multiple requests sequentially
	for i := 0; i < count; i++ {
		userAgent := userAgents[rand.Intn(len(userAgents))]
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Error creating request %d: %v\n", i+1, err)
			continue
		}
		req.Header.Set("User-Agent", userAgent)

		for attempt := 0; attempt < 3; attempt++ {
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("Error making request %d: %v\n", i+1, err)
				logFile.WriteString(fmt.Sprintf("Request %d: %v\n", i+1, err))
				time.Sleep(1 * time.Second)
				continue
			}

			func() {
				defer resp.Body.Close()
				fmt.Printf("Request %d: Received status %s\n", i+1, resp.Status)
			}()
			break
		}

		time.Sleep(time.Duration(delayBetweenRequests) * time.Second)
	}

	fmt.Println("All requests completed.")
}
