package command

import (
	"context"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/senchabot-opensource/monorepo/packages/gosenchabot/models"
)

func (c *commands) CmdsCommand(context context.Context, message twitch.PrivateMessage, commandName string, params []string) (*models.CommandResponse, error) {
	var cmdResp models.CommandResponse
	var commandListArr []string
	var commandListString string

	cmdResp.Channel = message.Channel

	commandList, err := c.service.GetCommandList(context, message.RoomID)
	if err != nil {
		return nil, err
	}

	for _, v := range commandList {
		commandListArr = append(commandListArr, v.CommandName)
	}

	commandListString = strings.Join(commandListArr, ", ")

	cmdResp.Message = message.Channel + "'s Channel Commands: " + commandListString
	return &cmdResp, nil
}
