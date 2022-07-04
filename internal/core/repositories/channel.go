package repositories

import (
	"context"
	"wynn-member-api/ent"
)

type ChannelRepository interface {
	GetChannels(ctx context.Context) ([]*ent.Channel, error)
}
