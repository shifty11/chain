package types

import (
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PoolKeeper interface {
	AssertPoolExists(ctx sdk.Context, poolId uint64) error
}

type StakersKeeper interface {
	DoesStakerExist(sdk.Context, string) bool
	DoesValaccountExist(sdk.Context, uint64, string) bool
	GetValaccountsFromStaker(ctx sdk.Context, stakerAddress string) (val []*stakerstypes.Valaccount)
	GetActiveStakers(sdk.Context) []string
	GetAllStakerAddressesOfPool(sdk.Context, uint64) []string
	GetPoolCount(sdk.Context, string) uint64
}
