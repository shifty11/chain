package util

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// Mint
	mintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	// Upgrade
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

type AccountKeeper interface {
	GetModuleAddress(string) sdk.AccAddress
}

type BankKeeper interface {
	GetBalance(sdk.Context, sdk.AccAddress, string) sdk.Coin
	GetSupply(sdk.Context, string) sdk.Coin
	SendCoins(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) error
	SendCoinsFromAccountToModule(sdk.Context, sdk.AccAddress, string, sdk.Coins) error
	SendCoinsFromModuleToAccount(sdk.Context, string, sdk.AccAddress, sdk.Coins) error
	SendCoinsFromModuleToModule(sdk.Context, string, string, sdk.Coins) error
}

type DistributionKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type MintKeeper interface {
	GetMinter(sdk.Context) mintTypes.Minter
	GetParams(sdk.Context) mintTypes.Params
}

type UpgradeKeeper interface {
	ScheduleUpgrade(sdk.Context, upgradeTypes.Plan) error
}
