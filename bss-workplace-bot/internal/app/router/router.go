package router

import (
	"context"
	bss_workplace_api "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-api"
	bss_workplace_facade "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-facade"
	"github.com/ozonmp/omp-bot/internal/app/commands/business"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, callback *tgbotapi.Message, commandPath path.CommandPath)
}

type Router struct {
	// bot
	bot               *tgbotapi.BotAPI
	businessCommander Commander
}

func NewRouter(bot *tgbotapi.BotAPI, apiClient bss_workplace_api.BssWorkplaceApiServiceClient, grpcFacadeClient bss_workplace_facade.BssFacadeEventsApiServiceClient) *Router {
	return &Router{
		bot:               bot,
		businessCommander: business.NewBusinessCommander(bot, apiClient, grpcFacadeClient),
	}
}

func (c *Router) HandleUpdate(ctx context.Context, update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(ctx, update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(ctx, update.Message)
	}
}

func (c *Router) handleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	switch callbackPath.Domain {
	case "business":
		c.businessCommander.HandleCallback(ctx, callback, callbackPath)
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", callbackPath.Domain)
	}
}

func (c *Router) handleMessage(ctx context.Context, msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		c.showCommandFormat(msg)

		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", msg.Command(), err)
		return
	}

	switch commandPath.Domain {
	case "business":
		c.businessCommander.HandleCommand(ctx, msg, commandPath)
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", commandPath.Domain)
	}
}

func (c *Router) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}__{domain}__{subdomain}")

	_, err := c.bot.Send(outputMsg)
	if err != nil {
		log.Printf("Router.showCommandFormat: error sending reply message to chat - %v", err)
	}
}
