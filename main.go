package main

import (
	"log"
	"os"

	"github.com/asticode/go-astisub"
)

const (
	SampleSubFile = "./samples/sample.srt"
	BatchLines    = 10
	TotalLines    = -1
)

func main() {
	s, err := astisub.OpenFile(SampleSubFile)
	if err != nil {
		log.Fatalf("failed to open subtitle file %s: %v", SampleSubFile, err)
	}

	APIKey := os.Getenv("API_KEY")
	if APIKey == "" {
		log.Fatal("Google translate API_KEY is empty or not set in environment variable")
	}

	translateClient := NewTranslateClient(APIKey)

	log.Printf("subtitle has %d lines", len(s.Items))

	inputs := []string{}
	total := 0
	totalTranslate := 0

	start := 0
	for i, item := range s.Items {
		inputs = addLines(item.Lines, inputs)
		total++

		if TotalLines > 0 && total >= TotalLines {
			break
		}

		if len(inputs) >= BatchLines {
			translated := translateClient.TranslateToJp(inputs)
			addTranslatedLines(s.Items, start, translated)
			totalTranslate += len(translated)

			start = i + 1
			inputs = []string{}
		}
	}

	if len(inputs) > 0 {
		translated := translateClient.TranslateToJp(inputs)
		addTranslatedLines(s.Items, start, translated)
		totalTranslate += len(translated)
	}

	log.Printf("total translated lines: %d", totalTranslate)

	s.Write("samples/sample.out.srt")
}

func addLines(lines []astisub.Line, inputs []string) []string {
	line := ""
	for _, l := range lines {
		if line != "" {
			line += "\n"
		}
		text := ""
		for _, it := range l.Items {
			text += it.Text
		}
		line += text
	}

	return append(inputs, line)
}

func addTranslatedLines(items []*astisub.Item, start int, translations []string) {
	for _, t := range translations {
		item := items[start]
		translatedLineItems := []astisub.LineItem{
			{
				Text: t,
			},
		}

		item.Lines = []astisub.Line{astisub.Line{Items: translatedLineItems}}

		start++
	}
}
