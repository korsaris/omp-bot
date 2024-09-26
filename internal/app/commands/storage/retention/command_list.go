package retention

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RetentionCommanderImpl) List(inputMessage *tgbotapi.Message) {
	// TODO: JSON IMPL + Inline Ketboard

	outputMsgText := "Here all the elements: \n\n"

	retents, err := c.retentionService.List(0, 0)
	for _, r := range retents {
		outputMsgText += fmt.Sprintf("%v\n", r)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	/*serializedData, _ := json.Marshal(CallbackListData{
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
	)*/

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("StorageRetentionCommander.List: error sending reply message to chat - %v", err)
	}
}
