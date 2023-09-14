package keeper

import (
	"fmt"

	"github.com/KYVENetwork/chain/util"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	// Pool
	"github.com/KYVENetwork/chain/x/pool/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storeTypes.StoreKey
		memKey   storeTypes.StoreKey

		authority string

		stakersKeeper types.StakersKeeper
		accountKeeper types.AccountKeeper
		bankKeeper    util.BankKeeper
		distrkeeper   util.DistributionKeeper
		mintKeeper    util.MintKeeper
		upgradeKeeper util.UpgradeKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storeTypes.StoreKey,
	memKey storeTypes.StoreKey,

	authority string,

	accountKeeper types.AccountKeeper,
	bankKeeper util.BankKeeper,
	distrKeeper util.DistributionKeeper,
	mintKeeper util.MintKeeper,
	upgradeKeeper util.UpgradeKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		authority: authority,

		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		distrkeeper:   distrKeeper,
		mintKeeper:    mintKeeper,
		upgradeKeeper: upgradeKeeper,
	}
}

func (k Keeper) EnsurePoolAccount(ctx sdk.Context, id uint64) {
	name := fmt.Sprintf("%s/%d", types.ModuleName, id)

	address := authTypes.NewModuleAddress(name)
	account := k.accountKeeper.GetAccount(ctx, address)

	if account == nil {
		// account doesn't exist, initialise a new module account.
		account = authTypes.NewEmptyModuleAccount(name)
	} else {
		// account exists, adjust it to a module account.
		baseAccount := authTypes.NewBaseAccount(address, nil, account.GetAccountNumber(), 0)

		account = authTypes.NewModuleAccount(baseAccount, name)
	}

	k.accountKeeper.SetAccount(ctx, account)
}

func (k *Keeper) SetStakersKeeper(stakersKeeper types.StakersKeeper) {
	k.stakersKeeper = stakersKeeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) StoreKey() storeTypes.StoreKey {
	return k.storeKey
}
