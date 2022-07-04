package repositories

import (
	"context"
	"wynn-member-api/ent"
	"wynn-member-api/ent/channel"
	"wynn-member-api/internal/core/repositories"
)

type channelRepository struct {
	db *ent.Client
}

func NewChannelRepository(pg *ent.Client) repositories.ChannelRepository {
	return &channelRepository{db: pg}
}

func (r channelRepository) GetChannels(ctx context.Context) ([]*ent.Channel, error) {
	return r.db.Channel.Query().Where(channel.StatusEQ("active")).All(ctx)
}
