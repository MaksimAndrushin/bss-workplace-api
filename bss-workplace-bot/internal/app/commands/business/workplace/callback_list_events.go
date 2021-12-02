package workplace

import (
	"context"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListEventsData struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

func (c *BusinessWorkplaceCommander) CallbackListEvents(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var parsedData = CallbackListEventsData{}

	if err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData); err != nil {
		description := fmt.Sprintf("Fail to unmarshal input data: " + callbackPath.CallbackData)
		c.processError(callback.Message.Chat.ID, description, "")
		return
	}

	c.processListEvents(ctx, parsedData.Offset, parsedData.Limit, callback.Message.Chat.ID)
}
