package tests

import (
	"context"
	"log"
	"testing"

	"golang.org/x/text/language"

	"cloud.google.com/go/translate"
	"google.golang.org/api/option"
)

const API_KEY = "AIzaSyAdb5Ocb1cTyLMIqOq69JHQBpKJm_rVeAw"

func TestAPI(t *testing.T) {
	ctx := context.Background()

	c, err := translate.NewClient(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		t.Error(err)
	}

	resp, err := c.Translate(ctx, []string{"hello world!"}, language.Japanese, &translate.Options{
		Source: language.English,
		Format: translate.Text,
		Model:  "nmt",
	})
	if err != nil {
		t.Error(err)
	}

	log.Printf("%+v", resp)
}
