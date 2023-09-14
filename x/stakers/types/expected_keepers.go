package types

import (
	poolTypes "github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

type PoolKeeper interface {
	AssertPoolExists(ctx sdk.Context, poolId uint64) error
	GetPoolWithError(ctx sdk.Context, poolId uint64) (poolTypes.Pool, error)
}

type UpgradeKeeper interface {
	ScheduleUpgrade(ctx sdk.Context, plan types.Plan) error
}

type DelegationKeeper interface {
	GetDelegationAmount(ctx sdk.Context, staker string) uint64
	GetDelegationAmountOfDelegator(ctx sdk.Context, stakerAddress string, delegatorAddress string) uint64
	GetStakersByDelegator(ctx sdk.Context, delegator string) []string
	Delegate(sdk.Context, string, string, uint64) error
}
