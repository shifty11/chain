package app

import (
	"time"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
	"cosmossdk.io/core/appconfig"
	stakersmodule "github.com/KYVENetwork/chain/pulsar/kyve/stakers/module/v1"
	"google.golang.org/protobuf/types/known/durationpb"

	// Auth
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	// Authz
	authzmodulev1 "cosmossdk.io/api/cosmos/authz/module/v1"
	"github.com/cosmos/cosmos-sdk/x/authz"
	// Bank
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	// Bundles
	bundlesmodule "github.com/KYVENetwork/chain/pulsar/kyve/bundles/module/v1"
	bundlestypes "github.com/KYVENetwork/chain/x/bundles/types"
	// Capability
	capabilitymodulev1 "cosmossdk.io/api/cosmos/capability/module/v1"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	// Consensus
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	// Crisis
	crisismodulev1 "cosmossdk.io/api/cosmos/crisis/module/v1"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	// Delegation
	delegationmodule "github.com/KYVENetwork/chain/pulsar/kyve/delegation/module/v1"
	delegationtypes "github.com/KYVENetwork/chain/x/delegation/types"
	// Distribution
	distrmodulev1 "cosmossdk.io/api/cosmos/distribution/module/v1"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	// Evidence
	evidencemodulev1 "cosmossdk.io/api/cosmos/evidence/module/v1"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	// FeeGrant
	feegrantmodulev1 "cosmossdk.io/api/cosmos/feegrant/module/v1"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	// Genutil
	genutilmodulev1 "cosmossdk.io/api/cosmos/genutil/module/v1"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	// Global
	globalmodule "github.com/KYVENetwork/chain/pulsar/kyve/global/module/v1"
	globaltypes "github.com/KYVENetwork/chain/x/global/types"
	// Governance
	govmodulev1 "cosmossdk.io/api/cosmos/gov/module/v1"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	// Group
	groupmodulev1 "cosmossdk.io/api/cosmos/group/module/v1"
	"github.com/cosmos/cosmos-sdk/x/group"
	// IBC Core
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	// IBC Fee
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	// IBC Transfer
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	// ICA
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	// Mint
	mintmodulev1 "cosmossdk.io/api/cosmos/mint/module/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	// Params
	paramsmodulev1 "cosmossdk.io/api/cosmos/params/module/v1"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	// Pool
	poolmodule "github.com/KYVENetwork/chain/pulsar/kyve/pool/module/v1"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	// Query
	querymodule "github.com/KYVENetwork/chain/pulsar/kyve/query/module/v1"
	querytypes "github.com/KYVENetwork/chain/x/query/types"
	// Slashing
	slashingmodulev1 "cosmossdk.io/api/cosmos/slashing/module/v1"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
	// Staking
	stakingmodulev1 "cosmossdk.io/api/cosmos/staking/module/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	// Team
	teammodule "github.com/KYVENetwork/chain/pulsar/kyve/team/module/v1"
	teamtypes "github.com/KYVENetwork/chain/x/team/types"
	// Upgrade
	upgrademodulev1 "cosmossdk.io/api/cosmos/upgrade/module/v1"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	// Vesting
	vestingmodulev1 "cosmossdk.io/api/cosmos/vesting/module/v1"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	// this line is used by starport scaffolding # stargate/app/moduleImport
)

var (
	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	genesisModuleOrder = []string{
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		consensustypes.ModuleName,

		// Kyve modules
		pooltypes.ModuleName,
		delegationtypes.ModuleName,
		bundlestypes.ModuleName,
		querytypes.ModuleName,
		globaltypes.ModuleName,
		teamtypes.ModuleName,
		stakerstypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	}

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
	beginBlockers = []string{
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		// NOTE: x/team must be run before x/distribution and after x/mint.
		teamtypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		ibcexported.ModuleName,
		authtypes.ModuleName,
		ibctransfertypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		consensustypes.ModuleName,
		// Kyve modules
		pooltypes.ModuleName,
		delegationtypes.ModuleName,
		bundlestypes.ModuleName,
		querytypes.ModuleName,
		globaltypes.ModuleName,
		stakerstypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/beginBlockers
	}

	endBlockers = []string{
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		consensustypes.ModuleName,
		// Kyve modules
		pooltypes.ModuleName,
		delegationtypes.ModuleName,
		bundlestypes.ModuleName,
		querytypes.ModuleName,
		globaltypes.ModuleName,
		teamtypes.ModuleName,
		stakerstypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/endBlockers
	}

	// module account permissions
	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: distrtypes.ModuleName},
		{Account: icatypes.ModuleName},
		{Account: minttypes.ModuleName, Permissions: []string{authtypes.Minter}},
		{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: govtypes.ModuleName, Permissions: []string{authtypes.Burner}},
		{Account: ibctransfertypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},

		// KYVE
		{Account: bundlestypes.ModuleName},
		{Account: delegationtypes.ModuleName},
		{Account: pooltypes.ModuleName},
		{Account: globaltypes.ModuleName, Permissions: []string{authtypes.Burner}},
		{Account: teamtypes.ModuleName},
		{Account: stakerstypes.ModuleName},
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}

	// blocked account addresses
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		distrtypes.ModuleName,
		minttypes.ModuleName,
		stakingtypes.BondedPoolName,
		stakingtypes.NotBondedPoolName,
		// We allow the following module accounts to receive funds:
		// govtypes.ModuleName
	}

	// AppConfig application configuration (used by depinject)
	AppConfig = appconfig.Compose(&appv1alpha1.Config{
		Modules: []*appv1alpha1.ModuleConfig{
			{
				Name: "runtime",
				Config: appconfig.WrapAny(&runtimev1alpha1.Module{
					AppName:       Name,
					BeginBlockers: beginBlockers,
					EndBlockers:   endBlockers,
					OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
						{
							ModuleName: authtypes.ModuleName,
							KvStoreKey: "acc",
						},
					},
					InitGenesis: genesisModuleOrder,
					// When ExportGenesis is not specified, the export genesis module order
					// is equal to the init genesis order
					// ExportGenesis: genesisModuleOrder,
					// Uncomment if you want to set a custom migration order here.
					// OrderMigrations: nil,
				}),
			},
			{
				Name: authtypes.ModuleName,
				Config: appconfig.WrapAny(&authmodulev1.Module{
					Bech32Prefix:             "kyve",
					ModuleAccountPermissions: moduleAccPerms,
					// By default modules authority is the governance module. This is configurable with the following:
					// Authority: "group", // A custom module authority can be set using a module name
					// Authority: "cosmos1cwwv22j5ca08ggdv9c2uky355k908694z577tv", // or a specific address
				}),
			},
			{
				Name:   vestingtypes.ModuleName,
				Config: appconfig.WrapAny(&vestingmodulev1.Module{}),
			},
			{
				Name: banktypes.ModuleName,
				Config: appconfig.WrapAny(&bankmodulev1.Module{
					BlockedModuleAccountsOverride: blockAccAddrs,
				}),
			},
			{
				Name:   stakingtypes.ModuleName,
				Config: appconfig.WrapAny(&stakingmodulev1.Module{}),
			},
			{
				Name:   slashingtypes.ModuleName,
				Config: appconfig.WrapAny(&slashingmodulev1.Module{}),
			},
			{
				Name:   paramstypes.ModuleName,
				Config: appconfig.WrapAny(&paramsmodulev1.Module{}),
			},
			{
				Name:   "tx",
				Config: appconfig.WrapAny(&txconfigv1.Config{}),
			},
			{
				Name:   genutiltypes.ModuleName,
				Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
			},
			{
				Name:   authz.ModuleName,
				Config: appconfig.WrapAny(&authzmodulev1.Module{}),
			},
			{
				Name:   upgradetypes.ModuleName,
				Config: appconfig.WrapAny(&upgrademodulev1.Module{}),
			},
			{
				Name:   distrtypes.ModuleName,
				Config: appconfig.WrapAny(&distrmodulev1.Module{}),
			},
			{
				Name: capabilitytypes.ModuleName,
				Config: appconfig.WrapAny(&capabilitymodulev1.Module{
					SealKeeper: true,
				}),
			},
			{
				Name:   evidencetypes.ModuleName,
				Config: appconfig.WrapAny(&evidencemodulev1.Module{}),
			},
			{
				Name:   minttypes.ModuleName,
				Config: appconfig.WrapAny(&mintmodulev1.Module{}),
			},
			{
				Name: group.ModuleName,
				Config: appconfig.WrapAny(&groupmodulev1.Module{
					MaxExecutionPeriod: durationpb.New(time.Second * 1209600),
					MaxMetadataLen:     255,
				}),
			},
			{
				Name:   feegrant.ModuleName,
				Config: appconfig.WrapAny(&feegrantmodulev1.Module{}),
			},
			{
				Name:   govtypes.ModuleName,
				Config: appconfig.WrapAny(&govmodulev1.Module{}),
			},
			{
				Name:   crisistypes.ModuleName,
				Config: appconfig.WrapAny(&crisismodulev1.Module{}),
			},
			{
				Name:   consensustypes.ModuleName,
				Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
			},
			// ----- KYVE Modules -----
			{
				Name:   bundlestypes.ModuleName,
				Config: appconfig.WrapAny(&bundlesmodule.Module{}),
			},
			{
				Name:   delegationtypes.ModuleName,
				Config: appconfig.WrapAny(&delegationmodule.Module{}),
			},
			{
				Name:   globaltypes.ModuleName,
				Config: appconfig.WrapAny(&globalmodule.Module{}),
			},
			{
				Name:   pooltypes.ModuleName,
				Config: appconfig.WrapAny(&poolmodule.Module{}),
			},
			{
				Name:   querytypes.ModuleName,
				Config: appconfig.WrapAny(&querymodule.Module{}),
			},
			{
				Name:   teamtypes.ModuleName,
				Config: appconfig.WrapAny(&teammodule.Module{}),
			},
			{
				Name:   stakerstypes.ModuleName,
				Config: appconfig.WrapAny(&stakersmodule.Module{}),
			},
			// this line is used by starport scaffolding # stargate/app/moduleConfig
		},
	})
)
