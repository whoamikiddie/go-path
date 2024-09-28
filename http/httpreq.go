package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a domain: ")

	domain, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	domain = strings.TrimSpace(domain)

	if !strings.HasPrefix(domain, "http://") && !strings.HasPrefix(domain, "https://") {
		domain = "http://" + domain
	}

	response, err := http.Get(domain)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	fmt.Printf("Response status code: %d\n", response.StatusCode)

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", response.StatusCode)
		return
	}

}
