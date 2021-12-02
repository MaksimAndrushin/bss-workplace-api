package workplace

import (
	"context"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *BusinessWorkplaceCommander) processListEvents(ctx context.Context, offset uint64, limit uint64, messageChatId int64) {
	events, total, _ := c.workplaceService.ListEvents(ctx, offset, limit)

	var outputMsgText = fmt.Sprintf("Events(offset - %d, page size - %d, total - %d): \n\n", offset, limit, total)

	for _, e := range events {
		outputMsgText += e.String()
		outputMsgText += "\n"
	}

	var msg = tgbotapi.NewMessage(messageChatId, outputMsgText)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(c.createTgListEventsButtonPanel(offset, uint64(len(events)), total)...))
	c.bot.Send(msg)
}

func (c *BusinessWorkplaceCommander) createTgListEventsButtonPanel(currOffset uint64, currPageSize uint64, total uint64) []tgbotapi.InlineKeyboardButton {
	var tgButtons = make([]tgbotapi.InlineKeyboardButton, 0, 4)

	// First page button
	tgButtons = append(tgButtons, tgbotapi.NewInlineKeyboardButtonData("First page", c.createListEventsButtonInfo(0, eventsPageSize).String()))

	// Previous page button
	var prevOffset int64 = int64(currOffset) - int64(currPageSize) - 1
	if prevOffset >= 0 {
		tgButtons = append(tgButtons, tgbotapi.NewInlineKeyboardButtonData("Prev page", c.createListEventsButtonInfo(uint64(prevOffset), currPageSize).String()))
	}

	// Next page button
	var nextOffset = int64(currOffset) + int64(currPageSize)
	if nextOffset <= int64(total-1) {
		tgButtons = append(tgButtons, tgbotapi.NewInlineKeyboardButtonData("Next page", c.createListEventsButtonInfo(uint64(nextOffset), currPageSize).String()))
	}

	// Last page button
	tgButtons = append(tgButtons, tgbotapi.NewInlineKeyboardButtonData("Last page", c.createListEventsButtonInfo(total-eventsPageSize, eventsPageSize).String()))

	return tgButtons
}

func (c *BusinessWorkplaceCommander) createListEventsButtonInfo(dataOffset uint64, limits uint64) path.CallbackPath {
	serializedData, _ := json.Marshal(CallbackListData{
		Offset: dataOffset,
		Limit:  limits,
	})

	var callbackPath = path.CallbackPath{
		Domain:       "business",
		Subdomain:    "workplace",
		CallbackName: "list_events",
		CallbackData: string(serializedData),
	}

	return callbackPath
}
