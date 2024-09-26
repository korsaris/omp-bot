package retention

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/storage/retention"
)

func (c *RetentionCommanderImpl) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	retentionID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	pos, err := c.retentionService.Create(*retention.NewRetention(retentionID))
	if err != nil {
		log.Printf("fail to create element with retentionID=%d: %v", retentionID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Element with retentionID:%v successfully created, position %v", retentionID, pos),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("RetentionCommanderImpl.New: error sending reply message to chat - %v", err)
	}
}
