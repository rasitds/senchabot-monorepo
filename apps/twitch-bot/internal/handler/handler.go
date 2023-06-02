package handler

import (
	"net/http"

	"github.com/senchabot-dev/monorepo/apps/twitch-bot/client"
	"github.com/senchabot-dev/monorepo/apps/twitch-bot/internal/service"
)

type Handler interface {
	InitBotEventHandlers()
	InitHttpHandlers(mux *http.ServeMux)
}

type Handlers struct {
	joinedChannelList []string
	client            *client.Clients
	service           *service.Services
}

func (s *Handlers) InitBotEventHandlers() {
	PrivateMessage(s.client, s.service)
	s.joinedChannelList = BotJoin(s.client, s.service)
}

func (s *Handlers) InitHttpHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		s.service.Webhook.BotJoin(s.client, s.joinedChannelList, w, r)
	})
}

func NewHandlers(client *client.Clients, service *service.Services) Handler {
	return &Handlers{
		client:  client,
		service: service,
	}
}
