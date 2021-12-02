package business

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	bss_workplace_api "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-api"
	bss_workplace_facade "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-facade"
	"github.com/ozonmp/omp-bot/internal/app/commands/business/workplace"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, message *tgbotapi.Message, commandPath path.CommandPath)
}

type BusinessCommander struct {
	bot                *tgbotapi.BotAPI
	workplaceCommander Commander
}

func NewBusinessCommander(bot *tgbotapi.BotAPI, apiClient bss_workplace_api.BssWorkplaceApiServiceClient, grpcFacadeClient bss_workplace_facade.BssFacadeEventsApiServiceClient) *BusinessCommander {
	return &BusinessCommander{
		bot: bot,
		// workplaceCommander
		workplaceCommander: workplace.NewWorkplaceCommander(bot, apiClient, grpcFacadeClient),
	}
}

func (c *BusinessCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "workplace":
		c.workplaceCommander.HandleCallback(ctx, callback, callbackPath)
	default:
		log.Printf("BusinessCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *BusinessCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "workplace":
		c.workplaceCommander.HandleCommand(ctx, msg, commandPath)
	default:
		log.Printf("BusinessCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
