package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/kavinaravind/go-genai/client"
)

// welcomeMessage is the welcome message displayed to the user when the chatbot starts.
const welcomeMessage = `Welcome!
You can start chatting by typing your questions or statements and pressing Enter.
Type 'fresh' to start a new chat session.
Type 'exit' to quit the chatbot.

ChatBot: Hello! Ask me anything.`

func main() {
	apiKey := flag.String("api-key", "", "API key for the Generative AI API")
	flag.Parse()

	if *apiKey == "" {
		log.Fatal("Please provide an API key with the -api-key flag")
	}

	ctx := context.Background()

	// Create a new Generative AI client
	client, err := client.NewGenAIClient(ctx, *apiKey)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Create a new reader and writer
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	// Start a new chat session
	chat := client.StartNewChatSession(ctx, writer)

	// Start chatting
	fmt.Println(welcomeMessage)
	for {
		fmt.Print("You: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		switch input {
		case "exit":
			fmt.Println("ChatBot: Goodbye!")
			os.Exit(0)
		case "fresh":
			chat = client.StartNewChatSession(ctx, writer)
			fmt.Println("ChatBot: Chat History Cleared.")
			continue
		default:
			fmt.Print("ChatBot: ")
			err = chat(genai.Text(input))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
