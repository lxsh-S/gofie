package src

import (
	"bufio"
	"os"
	"strings"
)

// This func will open a file and look for a specific label inside
func GetSystemKey(filePath, key string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return "Unkown error occured when opening the file"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// les loop through the file
	for scanner.Scan() {
		line := scanner.Text()

		// then les check if we got our metric
		if strings.Contains(line, key) {
			parts := strings.Split(line, ":") // we'll handel the data that is seperated by ":"
			if len(parts) > 1 {
				data := strings.TrimSpace(parts[1])
				return data
			}
		}

		// les also handel data seperated by "="
		if strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			if len(parts) > 1 {
				data := strings.TrimSpace(parts[1])
				return data
			}
		}
	}
	return "key not found" // we'll only reach this line when the scanner didnt find a key in the file
}
