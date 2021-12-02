package workplace

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var WorkplaceJsonExample string

//func init() {
//	workplaceExample, _ := json.Marshal(business.Workplace{ID: 1, Title: "Title Example", EmployeeID: 1, WorkplaceNumber: 1, OfficeID: 1})
//	WorkplaceJsonExample = string(workplaceExample)
//}

func (c *BusinessWorkplaceCommander) processError(tgChatID int64, description string, solution string) {
	log.Println(description)
	var msg = tgbotapi.NewMessage(tgChatID, description+"\n"+solution)
	c.bot.Send(msg)
}
