package retention

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RetentionCommanderImpl) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	retentionID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	ok, err := c.retentionService.Remove(retentionID)
	if err != nil {
		log.Printf("fail to delete element with retentionID=%d: %v", retentionID, err)
		return
	}

	var msg_text string
	if ok {
		msg_text = "Successfully deleted"
	} else {
		msg_text = "Not found"
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Element with retentionID=%v : %s", retentionID, msg_text),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("RetentionCommanderImpl.Delete: error sending reply message to chat - %v", err)
	}
}
