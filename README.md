# GenAI Chatbot

A simple chatbot that uses the Go SDK for Google Generative AI.

## Getting Started

These instructions will get you up and running on your local machine for development purposes.

## Prerequisites

- Go (currently running: go version go1.22.4 darwin/arm64)
- Make (currently running: GNU Make 3.81)

## Building

```shell
git clone https://github.com/kavinaravind/go-genai.git
cd go-genai
make build
```

## Google Generative AI

You can start by [Getting an API Key](https://ai.google.dev/gemini-api/docs/api-key) to use the Gemini API.

## Running

```shell
./chatbot --api-key <API_KEY> # after building the binary via make build
```

![chatbot](./assets/chatbot.mp4)
