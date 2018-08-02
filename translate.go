package main

import (
	"context"
	"log"

	"golang.org/x/text/language"

	"google.golang.org/api/option"

	"cloud.google.com/go/translate"
)

// TranslateClient wrapper of google translate api
type TranslateClient struct {
	APIKey string
	client *translate.Client
	ctx    context.Context
}

// NewTranslateClient create new instance of translate client
func NewTranslateClient(APIKey string) *TranslateClient {
	ctx := context.Background()
	gClient, err := translate.NewClient(ctx, option.WithAPIKey(APIKey))
	if err != nil {
		log.Printf("failed to create google translate client, %v", err)
		return nil
	}

	return &TranslateClient{APIKey, gClient, ctx}
}

// TranslateToJp translate input English strings into Japanese
func (c *TranslateClient) TranslateToJp(inputs []string) []string {
	result := []string{}

	if len(inputs) <= 0 {
		log.Println("translation input is empty, skip it")
		return result
	}

	log.Printf("translating %d inputs", len(inputs))

	resp, err := c.client.Translate(c.ctx, inputs, language.Japanese, &translate.Options{
		Source: language.SimplifiedChinese,
		Format: translate.Text,
		Model:  "nmt",
	})
	if err != nil {
		log.Printf("failed to get translate result from google api, %v", err)
		return result
	}

	log.Printf("got %d results from google api", len(resp))

	for _, t := range resp {
		result = append(result, t.Text)
	}

	return result
}
