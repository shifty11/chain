package keeper

import (
	"fmt"

	"github.com/KYVENetwork/chain/util"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/KYVENetwork/chain/x/query/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey
		// paramstore paramtypes.Subspace	//TODO(rapha): what is ps for?

		accountKeeper    util.AccountKeeper
		bankKeeper       util.BankKeeper
		distrkeeper      util.DistributionKeeper
		poolKeeper       types.PoolKeeper
		stakerKeeper     types.StakersKeeper
		delegationKeeper types.DelegationKeeper
		bundleKeeper     types.BundlesKeeper // TODO: rename to bundlesKeeper
		globalKeeper     types.GlobalKeeper
		govKeeper        types.GovKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	// ps paramtypes.Subspace,	//TODO(rapha): what is ps for?

	accountKeeper util.AccountKeeper,
	bankKeeper util.BankKeeper,
	distrkeeper util.DistributionKeeper,
	poolKeeper types.PoolKeeper,
	stakerKeeper types.StakersKeeper,
	delegationKeeper types.DelegationKeeper,
	// bundleKeeper types.BundlesKeeper,
	globalKeeper types.GlobalKeeper,
	govKeeper types.GovKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// paramstore: ps,

		accountKeeper:    accountKeeper,
		bankKeeper:       bankKeeper,
		distrkeeper:      distrkeeper,
		poolKeeper:       poolKeeper,
		stakerKeeper:     stakerKeeper,
		delegationKeeper: delegationKeeper,
		// bundleKeeper:     bundleKeeper,
		globalKeeper: globalKeeper,
		govKeeper:    govKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
