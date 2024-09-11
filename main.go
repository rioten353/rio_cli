package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email"`
	Address Address  `json:"address"`
	Hobbies []string `json:"hobbies"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <filename> <count='show, lines, words'>")
		return
	}

	filename := os.Args[1]

	count := os.Args[2]

	if count == "words" && isTextFile(filename) {
		wordCount, err := countWordsFromTextFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		for word, count := range wordCount {
			fmt.Printf("%s: %d\n", word, count)
		}
	} else if count == "words" && isJSONFile(filename) {
		wordCount, err := countWordsFromJSONFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		for word, count := range wordCount {
			fmt.Printf("%s: %d\n", word, count)
		}
	} else if count == "lines" {
		lineCount, err := countLinesFromFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		fmt.Println("Number of lines:", lineCount)
	} else if count == "show" {
		err := showContents(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

	} else {
		fmt.Println("Invalid count type. Only 'words' and 'lines' are supported.")
		return
	}
}

func isTextFile(filename string) bool {
	return strings.HasSuffix(filename, ".txt")
}

func isJSONFile(filename string) bool {
	return strings.HasSuffix(filename, ".json")
}

func countWordsFromTextFile(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	wordCount := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			wordCount[word]++
		}
	}

	return wordCount, scanner.Err()
}

func countWordsFromJSONFile(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var person Person
	err = json.NewDecoder(file).Decode(&person)
	if err != nil {
		return nil, err
	}

	wordCount := make(map[string]int)

	// Extract words from the struct fields
	wordCount[person.Name]++
	wordCount[person.Email]++
	wordCount[person.Address.Street]++
	wordCount[person.Address.City]++
	wordCount[person.Address.State]++
	wordCount[person.Address.Zip]++

	for _, hobby := range person.Hobbies {
		wordCount[hobby]++
	}

	return wordCount, nil
}

func countLinesFromFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount, scanner.Err()
}

func showContents(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}
