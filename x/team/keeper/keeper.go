package keeper

import (
	"fmt"

	"github.com/KYVENetwork/chain/util"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// Team
	"github.com/KYVENetwork/chain/x/team/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storeTypes.StoreKey

		accountKeeper util.AccountKeeper
		bankKeeper    util.BankKeeper
		mintKeeper    util.MintKeeper
		upgradeKeeper util.UpgradeKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storeTypes.StoreKey,
	accountKeeper util.AccountKeeper,
	bankKeeper util.BankKeeper,
	mintKeeper util.MintKeeper,
	upgradeKeeper util.UpgradeKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		mintKeeper:    mintKeeper,
		upgradeKeeper: upgradeKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) StoreKey() storeTypes.StoreKey {
	return k.storeKey
}
