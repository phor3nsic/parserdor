/*
Fast coded by chatGPT
*/

package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Regular expression to identify domain and subdomain
	domainRegex := regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)
	subdomainRegex := regexp.MustCompile(`^([a-zA-Z0-9-]+\.){2,}[a-zA-Z]{2,}$`)

	// Function to check if the input is a URL
	isURL := func(s string) bool {
		parsedURL, err := url.ParseRequestURI(s)
		return err == nil && (parsedURL.Scheme == "http" || parsedURL.Scheme == "https")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		ninput := strings.Replace(input, "*.", "", -1)

		// Determine the type of input
		switch {
		case isURL(ninput):
			fmt.Printf("[url] %s\n", ninput)
		case subdomainRegex.MatchString(ninput):
			fmt.Printf("[subdomain] %s\n", ninput)
		case domainRegex.MatchString(ninput):
			fmt.Printf("[domain] %s\n", ninput)
		default:
			fmt.Println("Invalid input")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
