package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	GetAccount(sdk.Context, sdk.AccAddress) types.AccountI
	SetAccount(sdk.Context, types.AccountI)
}

type StakersKeeper interface {
	LeavePool(ctx sdk.Context, staker string, poolId uint64)
	GetAllStakerAddressesOfPool(ctx sdk.Context, poolId uint64) (stakers []string)
}
