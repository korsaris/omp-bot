package retention

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RetentionCommanderImpl) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__storage__retention - help\n"+
			"/get__storage__retention {RetentionID} - get entry with RetentionID\n"+
			"/list__storage__retention - list all entries\n"+
			"/delete__storage__retention {RetentionID} - delete entry with RetentionID\n"+
			"/new__storage__retention {arglist} - create new entry with arglist if not exist\n"+
			"/edit__storage__retention {arglist} - edit exist entry with arglist",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("RetentionCommanderImpl.Help: error sending reply message to chat - %v", err)
	}
}
