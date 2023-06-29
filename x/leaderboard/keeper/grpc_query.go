package keeper

import (
	"github.com/dotneko/leaderboard/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
