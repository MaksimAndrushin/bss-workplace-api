package workplace

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const eventsPageSize uint64 = 5

func (c *BusinessWorkplaceCommander) ListEvents(ctx context.Context, inputMessage *tgbotapi.Message) {
	c.processListEvents(ctx, 0, eventsPageSize, inputMessage.Chat.ID)
}
