package ante

import (
	upgrade "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibctransfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v8/modules/core/02-client/types" //nolint:staticcheck
	ibcconn "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	affiliates "github.com/nemo-network/v4-chain/protocol/x/affiliates/types"
	blocktime "github.com/nemo-network/v4-chain/protocol/x/blocktime/types"
	bridge "github.com/nemo-network/v4-chain/protocol/x/bridge/types"
	clob "github.com/nemo-network/v4-chain/protocol/x/clob/types"
	delaymsg "github.com/nemo-network/v4-chain/protocol/x/delaymsg/types"
	feetiers "github.com/nemo-network/v4-chain/protocol/x/feetiers/types"
	govplus "github.com/nemo-network/v4-chain/protocol/x/govplus/types"
	listing "github.com/nemo-network/v4-chain/protocol/x/listing/types"
	perpetuals "github.com/nemo-network/v4-chain/protocol/x/perpetuals/types"
	prices "github.com/nemo-network/v4-chain/protocol/x/prices/types"
	ratelimit "github.com/nemo-network/v4-chain/protocol/x/ratelimit/types"
	revshare "github.com/nemo-network/v4-chain/protocol/x/revshare/types"
	rewards "github.com/nemo-network/v4-chain/protocol/x/rewards/types"
	sending "github.com/nemo-network/v4-chain/protocol/x/sending/types"
	stats "github.com/nemo-network/v4-chain/protocol/x/stats/types"
	vault "github.com/nemo-network/v4-chain/protocol/x/vault/types"
	vest "github.com/nemo-network/v4-chain/protocol/x/vest/types"
)

// IsInternalMsg returns true if the given msg is an internal message.
func IsInternalMsg(msg sdk.Msg) bool {
	switch msg.(type) {
	case
		// ------- CosmosSDK default modules
		// auth
		*auth.MsgUpdateParams,

		// bank
		*bank.MsgSetSendEnabled,
		*bank.MsgUpdateParams,

		// consensus
		*consensus.MsgUpdateParams,

		// crisis
		*crisis.MsgUpdateParams,

		// distribution
		*distribution.MsgCommunityPoolSpend,
		*distribution.MsgUpdateParams,

		// gov
		*gov.MsgExecLegacyContent,
		*gov.MsgUpdateParams,

		// slashing
		*slashing.MsgUpdateParams,

		// staking
		*staking.MsgUpdateParams,

		// upgrade
		*upgrade.MsgCancelUpgrade,
		*upgrade.MsgSoftwareUpgrade,

		// ------- Custom modules
		// blocktime
		*blocktime.MsgUpdateDowntimeParams,

		// bridge
		*bridge.MsgCompleteBridge,
		*bridge.MsgUpdateEventParams,
		*bridge.MsgUpdateProposeParams,
		*bridge.MsgUpdateSafetyParams,

		// clob
		*clob.MsgCreateClobPair,
		*clob.MsgUpdateBlockRateLimitConfiguration,
		*clob.MsgUpdateClobPair,
		*clob.MsgUpdateEquityTierLimitConfiguration,
		*clob.MsgUpdateLiquidationsConfig,

		// delaymsg
		*delaymsg.MsgDelayMessage,

		// feetiers
		*feetiers.MsgUpdatePerpetualFeeParams,

		// govplus
		*govplus.MsgSlashValidator,

		// listing
		*listing.MsgSetMarketsHardCap,
		*listing.MsgSetListingVaultDepositParams,

		// perpetuals
		*perpetuals.MsgCreatePerpetual,
		*perpetuals.MsgSetLiquidityTier,
		*perpetuals.MsgUpdateParams,
		*perpetuals.MsgUpdatePerpetualParams,

		// prices
		*prices.MsgCreateOracleMarket,
		*prices.MsgUpdateMarketParam,

		// ratelimit
		*ratelimit.MsgSetLimitParams,
		*ratelimit.MsgSetLimitParamsResponse,

		// revshare
		*revshare.MsgSetMarketMapperRevenueShare,
		*revshare.MsgSetMarketMapperRevShareDetailsForMarket,

		// rewards
		*rewards.MsgUpdateParams,

		// sending
		*sending.MsgSendFromModuleToAccount,

		// stats
		*stats.MsgUpdateParams,

		// vault
		*vault.MsgSetVaultParams,
		*vault.MsgUpdateDefaultQuotingParams,

		// vest
		*vest.MsgDeleteVestEntry,
		*vest.MsgSetVestEntry,

		// ibc
		*icahosttypes.MsgUpdateParams,
		*ibctransfer.MsgUpdateParams,
		*ibcclient.MsgUpdateParams,
		*ibcconn.MsgUpdateParams,

		// affiliates
		*affiliates.MsgUpdateAffiliateTiers:

		return true

	default:
		return false
	}
}
