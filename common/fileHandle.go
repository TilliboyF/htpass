package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadHtpasswdFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	users := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}
		username, hashedPassword := parts[0], parts[1]
		users[username] = hashedPassword
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func WriteHtpasswdFile(filePath string, users map[string]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for username, hashedPassword := range users {
		line := fmt.Sprintf("%s:%s\n", username, hashedPassword)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
