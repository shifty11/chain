package keeper

import (
	"fmt"

	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/query/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		accountKeeper    util.AccountKeeper
		bankKeeper       util.BankKeeper
		distrkeeper      util.DistributionKeeper
		poolKeeper       types.PoolKeeper
		stakerKeeper     types.StakersKeeper
		delegationKeeper types.DelegationKeeper
		bundlesKeeper    types.BundlesKeeper
		globalKeeper     types.GlobalKeeper
		govKeeper        types.GovKeeper
		fundersKeeper    types.FundersKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,

	accountKeeper util.AccountKeeper,
	bankKeeper util.BankKeeper,
	distrkeeper util.DistributionKeeper,
	poolKeeper types.PoolKeeper,
	stakerKeeper types.StakersKeeper,
	delegationKeeper types.DelegationKeeper,
	globalKeeper types.GlobalKeeper,
	govKeeper types.GovKeeper,
	fundersKeeper types.FundersKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper:    accountKeeper,
		bankKeeper:       bankKeeper,
		distrkeeper:      distrkeeper,
		poolKeeper:       poolKeeper,
		stakerKeeper:     stakerKeeper,
		delegationKeeper: delegationKeeper,
		globalKeeper:     globalKeeper,
		govKeeper:        govKeeper,
		fundersKeeper:    fundersKeeper,
	}
}

func (k *Keeper) SetBundlesKeeper(bundlesKeeper types.BundlesKeeper) {
	k.bundlesKeeper = bundlesKeeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
