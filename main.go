package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

type Response struct {
	Output []OutputItem `json:"output"`
}

type OutputItem struct {
	Content []ContentItem `json:"content"`
}

type ContentItem struct {
	Text string `json:"text"`
}

const (
	model = "gpt-3.5-turbo-0125"
	url   = "https://api.openai.com/v1/responses"
)

func main() {
	var (
		answer   string
		question string
	)

	huh.NewInput().
		Title("How may I help you, sir?").
		Value(&question).
		Run()

	findAnswer := func(ctx context.Context) error {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			return errors.New("OPENAI_API_KEY environment variable not set")
		}

		payload := map[string]interface{}{
			"model":        model,
			"instructions": "Make sure your response is valid Markdown.",
			"input":        question,
		}

		body, err := json.Marshal(payload)
		if err != nil {
			return errors.New("could not marshal payload")
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
		if err != nil {
			return errors.New("could not make request")
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+apiKey)

		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("could not make request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("got unexpected response %d", resp.StatusCode)
		}

		dec := json.NewDecoder(resp.Body)
		var rsp Response
		if err := dec.Decode(&rsp); err != nil {
			return fmt.Errorf("error decoding response: %w", err)
		}

		if len(rsp.Output) > 0 && len(rsp.Output[0].Content) > 0 {
			answer = rsp.Output[0].Content[0].Text
		}
		return nil
	}

	if err := spinner.New().Title("At once...").ActionWithErr(findAnswer).Run(); err != nil {
		color.Red("Got an error: %s.", err.Error())
		os.Exit(1)
	}

	if answer == "" {
		answer = "Could not get answer"
	}

	out, err := glamour.Render(answer, "dark")
	if err != nil {
		color.Red("Got an error: %s.", err.Error())
		os.Exit(1)
	}

	fmt.Println(out)
}
