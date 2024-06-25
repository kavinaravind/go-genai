package client

import (
	"bufio"
	"context"
	"fmt"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// The Gemini 1.5 models are versatile and work with both text-only and multimodal prompts.
const model = "gemini-1.5-flash"

// GenAIClient is a client for the Generative AI API.
type GenAIClient struct {
	client *genai.Client
}

// NewGenAIClient creates a new Generative AI client.
func NewGenAIClient(ctx context.Context, apiKey string) (*GenAIClient, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &GenAIClient{client: client}, nil
}

// StartNewChatSession starts a new chat session.
func (c *GenAIClient) StartNewChatSession(ctx context.Context, writer *bufio.Writer) func(parts ...genai.Part) error {
	// Creates a new instance of the named generative model
	model := c.client.GenerativeModel(model)

	// Start a chat session
	cs := model.StartChat()

	// Return a closure that can send messages and handle responses
	return func(parts ...genai.Part) error {
		// Retrieve a streaming request
		iter := cs.SendMessageStream(ctx, parts...)

		// Iterate over the responses
		for {
			resp, err := iter.Next()
			if err != nil {
				if err == iterator.Done {
					break
				}
				return err
			}
			printResponse(writer, resp)
		}

		return nil
	}
}

// Close closes the client
func (c *GenAIClient) Close() error {
	return c.client.Close()
}

// printResponse prints the response to the writer
func printResponse(writer *bufio.Writer, resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				response := fmt.Sprintf("%s", part)
				for _, char := range response {
					fmt.Fprintf(writer, "%c", char)
					writer.Flush()
					time.Sleep(15 * time.Millisecond)
				}
			}
		}
	}
}
