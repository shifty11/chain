package types

import (
	delegationTypes "github.com/KYVENetwork/chain/x/delegation/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PoolKeeper interface {
	AssertPoolExists(ctx sdk.Context, poolId uint64) error
	GetPoolWithError(ctx sdk.Context, poolId uint64) (pooltypes.Pool, error)
	GetPool(ctx sdk.Context, id uint64) (val pooltypes.Pool, found bool)

	IncrementBundleInformation(ctx sdk.Context, poolId uint64, currentHeight uint64, currentKey string, currentValue string)

	GetAllPools(ctx sdk.Context) (list []pooltypes.Pool)
	GetProtocolInflationShare(ctx sdk.Context) (res sdk.Dec)
	ChargeInflationPool(ctx sdk.Context, poolId uint64) (payout uint64, err error)
}

type StakersKeeper interface {
	GetAllStakerAddressesOfPool(ctx sdk.Context, poolId uint64) (stakers []string)
	GetCommission(ctx sdk.Context, stakerAddress string) sdk.Dec
	IncreaseStakerCommissionRewards(ctx sdk.Context, address string, amount uint64) error
	AssertValaccountAuthorized(ctx sdk.Context, poolId uint64, stakerAddress string, valaddress string) error

	DoesStakerExist(ctx sdk.Context, staker string) bool
	DoesValaccountExist(ctx sdk.Context, poolId uint64, stakerAddress string) bool

	LeavePool(ctx sdk.Context, staker string, poolId uint64)

	IncrementPoints(ctx sdk.Context, poolId uint64, stakerAddress string) (newPoints uint64)
	ResetPoints(ctx sdk.Context, poolId uint64, stakerAddress string) (previousPoints uint64)
}

type DelegationKeeper interface {
	GetDelegationAmount(ctx sdk.Context, staker string) uint64
	GetDelegationOfPool(ctx sdk.Context, poolId uint64) uint64
	GetTotalAndHighestDelegationOfPool(ctx sdk.Context, poolId uint64) (uint64, uint64)
	PayoutRewards(ctx sdk.Context, staker string, amount uint64, payerModuleName string) error
	SlashDelegators(ctx sdk.Context, poolId uint64, staker string, slashType delegationTypes.SlashType)
}

type FundersKeeper interface {
	ChargeFundersOfPool(ctx sdk.Context, poolId uint64) (payout uint64, err error)
}

type TeamKeeper interface {
	GetTeamBlockProvision(ctx sdk.Context) int64
}
