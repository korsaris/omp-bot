package retention

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *RetentionCommanderImpl) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	retents, err := c.retentionService.List(0, 0)
	for _, r := range retents {
		outputMsgText += fmt.Sprintf("%d", r.RetentionID)
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "storage",
		Subdomain:    "retention",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("StorageRetentionCommander.List: error sending reply message to chat - %v", err)
	}
}
