package src

import (
	"bufio"
	"os"
	"strings"
)

// This func will open a file and look for a specific label inside
func GetSystemKey(filePath, key string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return " ", err
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
				return data, nil
			}
		}

		// les also handel data seperated by "="
		if strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			if len(parts) > 1 {
				data := strings.TrimSpace(parts[1])
				return data, nil
			}
		}
	}
	return "done", nil
}
