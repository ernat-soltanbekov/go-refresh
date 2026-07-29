package main

import (
	"fmt"
	"go-refresh/internal/ai"
	"go-refresh/internal/pipeline"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Используйте: go run . <input_file> <output_file> [--lang] [--keywords]")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	langFlag := false
	keywordsFlag := false

	for i := 3; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch arg {
		case "--lang":
			langFlag = true
		case "--keywords":
			keywordsFlag = true
		default:
			fmt.Fprintf(os.Stderr, "Ошибка: неизвестный флаг или дополнительный аргумент '%s'\n", arg)
			fmt.Fprintln(os.Stderr, "Используйте: go run . <input_file> <output_file> [--lang] [--keywords]")
			os.Exit(1)
		}
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	processedText := pipeline.Process(string(content))

	err = os.WriteFile(outputFile, []byte(processedText), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}

	if langFlag {
		fmt.Println(ai.LangDetect(processedText))
	}
	if keywordsFlag {
		fmt.Println(ai.ExtractKeywords(processedText))
	}
}
