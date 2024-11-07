package slackService

import (
	"context"

	"github.com/slack-go/slack"
)

type SlackService struct {
	sh SlackBotHandler
}

func NewSlackService(sh SlackBotHandler) *SlackService {
	return &SlackService{
		sh: sh,
	}
}

func (ss *SlackService) SendMessage(ctx context.Context,message string) error {
	_, _, err := ss.sh.SlackClient().PostMessage("api", slack.MsgOptionText(message, false))
	if err != nil {
		return err
	}
	return nil
}