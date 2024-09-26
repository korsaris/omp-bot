package retention

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/storage/retention"
)

func (c *RetentionCommanderImpl) Edit(inputMessage *tgbotapi.Message) {
	// TODO: JSON
	args := inputMessage.CommandArguments()

	retentionID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	err = c.retentionService.Update(retentionID, *retention.NewRetention(retentionID))
	if err != nil {
		log.Printf("fail to update element with retentionID=%d: %v", retentionID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Element with retentionID:%v successfully updated", retentionID),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("RetentionCommanderImpl.Edit: error sending reply message to chat - %v", err)
	}
}
