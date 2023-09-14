package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BankKeeper interface {
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	GetAllBalances(sdk.Context, sdk.AccAddress) sdk.Coins
}
