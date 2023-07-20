package command

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/senchabot-opensource/monorepo/apps/twitch-bot/internal/models"
	"github.com/senchabot-opensource/monorepo/packages/gosenchabot"
)

func (c *commands) InviteCommand(context context.Context, message twitch.PrivateMessage, commandName string, params []string) (*models.CommandResponse, error) {
	var cmdResp models.CommandResponse
	cmdResp.Channel = message.Channel

	if message.Channel != "senchabot" {
		return nil, errors.New("command did not executed in senchabot")
	}

	if len(params) < 1 {
		cmdResp.Message = "!invite [your_channel_name]"
		return &cmdResp, nil
	}

	var channelName = message.User.Name
	if strings.ToLower(params[0]) != channelName {
		cmdResp.Message = "You need to specify your channel name. !invite " + channelName
		return &cmdResp, nil
	}

	var twitchChannelId = message.User.ID
	alreadyJoined, err := c.service.CreateTwitchChannel(context, twitchChannelId, channelName, nil)
	if err != nil {
		return nil, errors.New("(CreateTwitchChannel) Error: " + err.Error())
	}

	if alreadyJoined {
		return nil, errors.New("already joined")
	}

	fmt.Println("TRYING TO JOIN TWITCH CHANNEL `" + channelName + "`")
	c.client.Twitch.Join(channelName)
	optionalCommands := gosenchabot.GetOptionalCommands()
	for _, command := range optionalCommands {
		_, err := c.service.CreateBotCommand(context, command.CommandName, command.CommandContent, twitchChannelId, "Senchabot")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	cmdResp.Message = "I joined your Twitch channel, sweetie"
	return &cmdResp, nil
}
