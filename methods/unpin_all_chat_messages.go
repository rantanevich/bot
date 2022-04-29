package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type UnpinAllChatMessagesParams struct {
	ChatID any `json:"chat_id" rules:"required,type:string|int"`
}

// UnpinAllChatMessages https://core.telegram.org/bots/api#unpinallchatmessages
func UnpinAllChatMessages(ctx context.Context, b *bot.Bot, params *UnpinAllChatMessagesParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "unpinAllChatMessages", params, &result)

	return result, err
}
