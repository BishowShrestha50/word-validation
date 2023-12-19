package service

import (
	"bufio"
	"os"
	"strings"
)

func (s *Service) GetAllWords(filePath string) (map[string]bool, error) {
	wordSet := make(map[string]bool)
	readFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		wordSet[strings.ToLower(fileScanner.Text())] = true
	}
	return wordSet, nil
}
