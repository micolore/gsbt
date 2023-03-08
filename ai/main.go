package main

import (
	"context"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type Model struct {
	Name     string
	MaxToken int
}

func main() {
	c := gogpt.NewClient("your token")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
