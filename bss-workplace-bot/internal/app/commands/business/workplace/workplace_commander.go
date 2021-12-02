package workplace

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	bss_workplace_api "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-api"
	bss_workplace_facade "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-facade"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/business/workplace"
	"log"
)

type WorkplaceCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(ctx context.Context, inputMsg *tgbotapi.Message)
	List(ctx context.Context, inputMsg *tgbotapi.Message)
	Delete(ctx context.Context, inputMsg *tgbotapi.Message)
	New(ctx context.Context, inputMsg *tgbotapi.Message)
	Edit(ctx context.Context, inputMsg *tgbotapi.Message)
	ListEvents(ctx context.Context, inputMsg *tgbotapi.Message)
}

type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, message *tgbotapi.Message, commandPath path.CommandPath)
}

type BusinessWorkplaceCommander struct {
	bot              *tgbotapi.BotAPI
	workplaceService *service.GrpcWorkplaceService
}

func NewWorkplaceCommander(bot *tgbotapi.BotAPI, apiClient bss_workplace_api.BssWorkplaceApiServiceClient, grpcFacadeClient bss_workplace_facade.BssFacadeEventsApiServiceClient) *BusinessWorkplaceCommander {
	var workplaceService = service.NewGrpcWorkplaceService(apiClient, grpcFacadeClient)

	return &BusinessWorkplaceCommander{
		bot:              bot,
		workplaceService: workplaceService,
	}
}

func (c *BusinessWorkplaceCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(ctx, callback, callbackPath)
	case "list_events":
		c.CallbackListEvents(ctx, callback, callbackPath)
	default:
		log.Printf("BusinessWorkplaceCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *BusinessWorkplaceCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(ctx, msg)
	case "events":
		c.ListEvents(ctx, msg)
	case "get":
		c.Get(ctx, msg)
	case "delete":
		c.Delete(ctx, msg)
	case "new":
		c.New(ctx, msg)
	case "edit":
		c.Edit(ctx, msg)
	default:
		c.Default(msg)
	}
}
