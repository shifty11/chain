package keeper

import (
	"fmt"

	"github.com/KYVENetwork/chain/util"
	"github.com/cometbft/cometbft/libs/log"

	"github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		authority string

		accountKeeper    util.AccountKeeper
		bankKeeper       util.BankKeeper
		distrkeeper      util.DistributionKeeper
		poolKeeper       types.PoolKeeper
		upgradeKeeper    types.UpgradeKeeper
		delegationKeeper types.DelegationKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	memKey storetypes.StoreKey,

	authority string,

	accountKeeper util.AccountKeeper,
	bankKeeper util.BankKeeper,
	distrkeeper util.DistributionKeeper,
	poolKeeper types.PoolKeeper,
	upgradeKeeper types.UpgradeKeeper,
	delegationKeeper types.DelegationKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		authority: authority,

		accountKeeper:    accountKeeper,
		bankKeeper:       bankKeeper,
		distrkeeper:      distrkeeper,
		poolKeeper:       poolKeeper,
		upgradeKeeper:    upgradeKeeper,
		delegationKeeper: delegationKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) StoreKey() storetypes.StoreKey {
	return k.storeKey
}
