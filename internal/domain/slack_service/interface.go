package slackService

import "github.com/slack-go/slack"

type SlackBotHandler interface {
	SlackClient() *slack.Client
}