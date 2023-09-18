package types

import (
	bundlesTypes "github.com/KYVENetwork/chain/x/bundles/types"
	delegationTypes "github.com/KYVENetwork/chain/x/delegation/types"
	globalTypes "github.com/KYVENetwork/chain/x/global/types"
	poolTypes "github.com/KYVENetwork/chain/x/pool/types"
	stakersTypes "github.com/KYVENetwork/chain/x/stakers/types"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

type GovKeeper interface {
	GetParams(sdk.Context) govTypes.Params
}

// TODO(rapha): do we need all of this methods?
type BundlesKeeper interface {
	AssertCanVote(sdk.Context, uint64, string, string, string) error
	AssertCanPropose(sdk.Context, uint64, string, string, uint64) error
	//GetBundleProposal(sdk.Context, uint64) (bundlesTypes.BundleProposal, bool)
	GetBundleVersionMap(sdk.Context) bundlesTypes.BundleVersionMap
	//GetFinalizedBundle(sdk.Context, uint64, uint64) (bundlesTypes.FinalizedBundle, bool)
	//GetFinalizedBundleByHeight(sdk.Context, uint64, uint64) (FinalizedBundle, bool)
	GetFinalizedBundleByIndex(sdk.Context, uint64, uint64) (FinalizedBundle, bool)
	//GetPaginatedFinalizedBundleQuery(sdk.Context, *query.PageRequest, uint64) ([]FinalizedBundle, *query.PageResponse, error)
	//GetParams(sdk.Context) bundlesTypes.Params
	//GetVoteDistribution(sdk.Context, uint64) bundlesTypes.VoteDistribution

	GetBundleProposal(sdk.Context, uint64) (bundlesTypes.BundleProposal, bool)
	GetFinalizedBundle(sdk.Context, uint64, uint64) (bundlesTypes.FinalizedBundle, bool)
	//GetFinalizedBundleByHeight(sdk.Context, uint64, uint64) (bundlesTypes.FinalizedBundle, bool)
	GetPaginatedFinalizedBundleQuery(sdk.Context, *query.PageRequest, uint64) ([]FinalizedBundle, *query.PageResponse, error)
	GetParams(sdk.Context) bundlesTypes.Params
	GetVoteDistribution(sdk.Context, uint64) bundlesTypes.VoteDistribution
}

type DelegationKeeper interface {
	GetDelegationAmount(sdk.Context, string) uint64
	GetDelegationAmountOfDelegator(sdk.Context, string, string) uint64
	GetDelegationOfPool(sdk.Context, uint64) uint64
	GetOutstandingRewards(sdk.Context, string, string) uint64
	GetPaginatedActiveStakersByDelegation(sdk.Context, *query.PageRequest, func(string, bool) bool) (*query.PageResponse, error)
	GetPaginatedActiveStakersByPoolCountAndDelegation(sdk.Context, *query.PageRequest) ([]string, *query.PageResponse, error)
	GetPaginatedInactiveStakersByDelegation(sdk.Context, *query.PageRequest, func(string, bool) bool) (*query.PageResponse, error)
	GetPaginatedStakersByDelegation(sdk.Context, *query.PageRequest, func(string, bool) bool) (*query.PageResponse, error)
	GetRedelegationCooldown(sdk.Context) uint64
	GetRedelegationCooldownEntries(sdk.Context, string) []uint64
	GetRedelegationMaxAmount(sdk.Context) uint64
	GetAllUnbondingDelegationQueueEntriesOfDelegator(sdk.Context, string) []delegationTypes.UndelegationQueueEntry
	GetDelegationData(sdk.Context, string) (delegationTypes.DelegationData, bool)
	GetUndelegationQueueEntry(sdk.Context, uint64) (delegationTypes.UndelegationQueueEntry, bool)
	GetParams(sdk.Context) delegationTypes.Params
	StoreKey() storeTypes.StoreKey
}

type GlobalKeeper interface {
	GetParams(sdk.Context) globalTypes.Params
}

type PoolKeeper interface {
	GetPaginatedPoolsQuery(sdk.Context, *query.PageRequest, string, string, bool, uint32) ([]poolTypes.Pool, *query.PageResponse, error)
	GetAllPools(sdk.Context) []poolTypes.Pool
	GetPool(sdk.Context, uint64) (poolTypes.Pool, bool)
	GetPoolWithError(sdk.Context, uint64) (poolTypes.Pool, error)
}

type StakersKeeper interface {
	DoesStakerExist(ctx sdk.Context, staker string) bool
	GetAllStakerAddressesOfPool(ctx sdk.Context, poolId uint64) (stakers []string)
	GetAllValaccountsOfPool(sdk.Context, uint64) []*stakersTypes.Valaccount
	GetCommissionChangeEntryByIndex2(sdk.Context, string) (stakersTypes.CommissionChangeEntry, bool)
	GetParams(sdk.Context) stakersTypes.Params
	GetStaker(sdk.Context, string) (stakersTypes.Staker, bool)
	GetValaccountsFromStaker(sdk.Context, string) []*stakersTypes.Valaccount
}
