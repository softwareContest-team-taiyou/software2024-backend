package infrastructure

import (
	"log"
	"os"

	"github.com/slack-go/slack"
)
type SlackBotHandler struct {
	Client *slack.Client
}
func NewSlackBotHandler() *SlackBotHandler {
	token := os.Getenv("SLACK_BOT_TOKEN") // Ensure your SLACK_BOT_TOKEN environment variable is set.
	if token == "" {
		log.Fatal("SLACK_BOT_TOKEN is required but not set")
	}

	client := slack.New(token, slack.OptionDebug(os.Getenv("ENV") == "development"))

	return &SlackBotHandler{Client: client}
}

func (h *SlackBotHandler) SlackClient() *slack.Client {
	return h.Client
}