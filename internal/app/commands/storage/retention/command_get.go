package retention

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RetentionCommanderImpl) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	retentionID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	retent, err := c.retentionService.Describe(retentionID)
	if err != nil {
		log.Printf("fail to get element with retentionID=%d: %v", retentionID, err)
		return
	}

	// TODO: Retention Stringer or Json
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("retentionID:%v", retent.RetentionID),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("RetentionCommanderImpl.Get: error sending reply message to chat - %v", err)
	}
}
