# watson

Tired of switching from terminal to GPT on browser? `watson` is your personal assistant from your terminal.

Using GPT under the hood, `watson` makes sure you're staying focused in the terminal.

![GIF Example](https://github.com/zuzuleinen/watson/blob/main/showcase.gif)

## Requirements

- Go 1.24+
- A valid OpenAI API key

## Installation

Add your `OPENAI_API_KEY` in your environment.

```shell
export OPENAI_API_KEY=your_openai_api_key
```

Run it from the source code:

```shell
   git clone https://github.com/zuzuleinen/watson.git
   cd watson
   go run .
```

Or `go install` and use the generated binary.

## Info

`watson` is using `gpt-3.5-turbo-0125` model.

## Credits

Created using the amazing tools from [Charmbracelet](https://charm.sh/).