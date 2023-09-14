package bundles

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	"github.com/KYVENetwork/chain/util"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/cometbft/cometbft/abci/types"

	modulev1 "github.com/KYVENetwork/chain/pulsar/kyve/bundles/module/v1"
	"github.com/KYVENetwork/chain/x/bundles/client/cli"
	"github.com/KYVENetwork/chain/x/bundles/keeper"
	"github.com/KYVENetwork/chain/x/bundles/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	// this line is used by starport scaffolding # 1
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the name of the module as a string
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage. The default GenesisState need to be defined by the module developer and is primarily used for testing
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root Tx command for the module. The subcommands of this root command are used by end-users to generate new transactions containing messages defined in the module
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the root query command for the module. The subcommands of this root command are used by end-users to generate new queries to the subset of the state defined by the module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	bankKeeper    util.BankKeeper
	mintKeeper    util.MintKeeper
	upgradeKeeper util.UpgradeKeeper
	poolKeeper    types.PoolKeeper
	teamKeeper    types.TeamKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	bankKeeper util.BankKeeper,
	mintKeeper util.MintKeeper,
	upgradeKeeper util.UpgradeKeeper,
	poolKeeper types.PoolKeeper,
	teamKeeper types.TeamKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		bankKeeper:     bankKeeper,
		mintKeeper:     mintKeeper,
		upgradeKeeper:  upgradeKeeper,
		poolKeeper:     poolKeeper,
		teamKeeper:     teamKeeper,
	}
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value, the InvariantRegistry triggers appropriate logic (most often the chain will be halted)
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion is a sequence number for state-breaking change of the module. It should be incremented on each consensus-breaking change introduced by the module. To avoid wrong/empty versions, the initial version should be set to 1
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock contains the logic that is automatically triggered at the beginning of each block
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	am.keeper.InitMemStore(ctx)
	SplitInflation(ctx, am.keeper, am.bankKeeper, am.mintKeeper, am.poolKeeper, am.teamKeeper, am.upgradeKeeper)
}

// EndBlock contains the logic that is automatically triggered at the end of each block
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	am.keeper.HandleUploadTimeout(sdk.WrapSDKContext(ctx))
	return []abci.ValidatorUpdate{}
}

// App Wiring Setup
func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type BundlesInputs struct {
	depinject.In

	Config *modulev1.Module
	Cdc    codec.Codec
	Key    *storeTypes.KVStoreKey
	MemKey *storeTypes.MemoryStoreKey

	AccountKeeper      util.AccountKeeper
	BankKeeper         util.BankKeeper
	DelegationKeeper   types.DelegationKeeper
	DistributionKeeper util.DistributionKeeper
	PoolKeeper         types.PoolKeeper
	StakersKeeper      types.StakersKeeper
	UpgradeKeeper      util.UpgradeKeeper
	MintKeeper         util.MintKeeper
	TeamKeeper         types.TeamKeeper
}

type BundlesOutpus struct {
	depinject.Out

	BundlesKeeper *keeper.Keeper
	Module        appmodule.AppModule
}

func ProvideModule(in BundlesInputs) BundlesOutpus {
	authority := authTypes.NewModuleAddress(govTypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authTypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	bundlesKeeper := keeper.NewKeeper(
		in.Cdc,
		in.Key,
		in.MemKey,
		authority.String(),
		in.AccountKeeper,
		in.BankKeeper,
		in.DistributionKeeper,
		in.PoolKeeper,
		in.StakersKeeper,
		in.DelegationKeeper,
	)
	m := NewAppModule(in.Cdc, *bundlesKeeper, in.BankKeeper, in.MintKeeper, in.UpgradeKeeper, in.PoolKeeper, in.TeamKeeper)

	return BundlesOutpus{BundlesKeeper: bundlesKeeper, Module: m}
}
