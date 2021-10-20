package retention

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RetentionCommanderImpl) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	retent_list, err := c.retentionService.List(idx, 1)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	retent := retent_list[0]
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("retentionID:%v", retent.RetentionID),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("RetentionCommanderImpl.Get: error sending reply message to chat - %v", err)
	}
}
