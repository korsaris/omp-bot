package retention

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/storage/retention"
)

type RetentionCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type RetentionCommanderImpl struct {
	bot              *tgbotapi.BotAPI
	retentionService retention.RetentionService
}

func NewRetentionCommanderImpl(
	bot *tgbotapi.BotAPI,
) *RetentionCommanderImpl {
	return &RetentionCommanderImpl{
		bot:              bot,
		retentionService: retention.NewDummyRetentionService(),
	}
}

func (c *RetentionCommanderImpl) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
	//	c.CallbackList(callback, callbackPath)
	default:
		log.Printf("RetentionCommanderImpl.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *RetentionCommanderImpl) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
