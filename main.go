package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	gpt3 "github.com/PullRequestInc/go-gpt3"
)

func Process(prompt string, client gpt3.Client, ctx context.Context, silent bool) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			prompt,
		},
		MaxTokens:   gpt3.IntPtr(2000),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		if !silent {
			fmt.Print(resp.Choices[0].Text)
		}
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if !silent {
		fmt.Printf("\n")
	}
}

func main() {
	openAIKey := os.Getenv("OPENAI_API_KEY")
	if openAIKey == "" {
		panic("Make sure you set your openAI API key in OPENAI_API_KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(openAIKey)

	intro := "My name is Michael. I'm a software engineer. During this entire session, write all your answers using markdown."
	Process(intro, client, ctx, true)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("gpt> ")
		if !scanner.Scan() {
			break
		}
		prompt := scanner.Text()
		Process(prompt, client, ctx, false)
	}
}
