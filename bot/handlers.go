package bot

import (
	"github.com/gofrs/uuid"
	. "github.com/traPtitech/traQ/model"
	"github.com/traPtitech/traQ/utils/message"
	"go.uber.org/zap"
)

func (p *Processor) pingHandler(bot *Bot) {
	payload := pingPayload{
		basePayload: makeBasePayload(),
	}

	buf, release, err := p.makePayloadJSON(&payload)
	if err != nil {
		p.logger.Error("unexpected json encode error", zap.Error(err))
		return
	}
	defer release()

	if p.sendEvent(bot, Ping, buf) {
		// OK
		if err := p.repo.ChangeBotState(bot.ID, BotActive); err != nil {
			p.logger.Error("failed to ChangeBotState", zap.Error(err))
		}
	} else {
		// NG
		if err := p.repo.ChangeBotState(bot.ID, BotPaused); err != nil {
			p.logger.Error("failed to ChangeBotState", zap.Error(err))
		}
	}
}

func (p *Processor) joinedAndLeftHandler(botID, channelID uuid.UUID, ev BotEvent) {
	bot, err := p.repo.GetBotByID(botID)
	if err != nil {
		p.logger.Error("failed to GetBotByID", zap.Error(err), zap.Stringer("id", botID))
		return
	}

	if filterBot(p, bot, stateFilter(BotActive)) {
		return
	}

	payload := joinAndLeftPayload{
		basePayload: makeBasePayload(),
		ChannelId:   channelID,
	}

	buf, release, err := p.makePayloadJSON(&payload)
	if err != nil {
		p.logger.Error("unexpected json encode error", zap.Error(err))
		return
	}
	defer release()

	p.sendEvent(bot, ev, buf)
}

func (p *Processor) createMessageHandler(message *Message, embedded []*message.EmbeddedInfo, plain string) {
	bots, err := p.repo.GetBotsByChannel(message.ChannelID)
	if err != nil {
		p.logger.Error("failed to GetBotsByChannel", zap.Error(err))
		return
	}
	bots = filterBots(p, bots, stateFilter(BotActive), eventFilter(MessageCreated), botUserIDNotEqualsFilter(message.UserID))
	if len(bots) == 0 {
		return
	}

	payload := messageCreatedPayload{
		basePayload: makeBasePayload(),
		Message: messagePayload{
			ID:        message.ID,
			UserID:    message.UserID,
			ChannelID: message.ChannelID,
			Text:      message.Text,
			PlainText: plain,
			Embedded:  embedded,
			CreatedAt: message.CreatedAt,
			UpdatedAt: message.UpdatedAt,
		},
	}

	buf, release, err := p.makePayloadJSON(&payload)
	if err != nil {
		p.logger.Error("unexpected json encode error", zap.Error(err))
		return
	}
	defer release()

	for _, bot := range bots {
		p.sendEvent(bot, MessageCreated, buf)
	}
}
