package workplace

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const pageSize uint64 = 5

func (c *BusinessWorkplaceCommander) List(ctx context.Context, inputMessage *tgbotapi.Message) {
	c.processList(ctx, 0, pageSize, inputMessage.Chat.ID)
}
