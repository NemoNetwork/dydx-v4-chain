package msgs

import (
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/nemo-network/v4-chain/protocol/testutil/constants"
	bridgetypes "github.com/nemo-network/v4-chain/protocol/x/bridge/types"
	clobtypes "github.com/nemo-network/v4-chain/protocol/x/clob/types"
	perptypes "github.com/nemo-network/v4-chain/protocol/x/perpetuals/types"
	pricestypes "github.com/nemo-network/v4-chain/protocol/x/prices/types"
)

var (
	// AppInjectedMsgSamples are msgs that are injected into the block by the proposing validator.
	// These messages are reserved for proposing validator's use only.
	AppInjectedMsgSamples = map[string]sdk.Msg{
		// bridge
		"/nemo_network.bridge.MsgAcknowledgeBridges": &bridgetypes.MsgAcknowledgeBridges{
			Events: []bridgetypes.BridgeEvent{
				{
					Id: 0,
					Coin: sdk.NewCoin(
						"bridge-token",
						sdkmath.NewIntFromUint64(1234),
					),
					Address: constants.Alice_Num0.Owner,
				},
			},
		},
		"/nemo_network.bridge.MsgAcknowledgeBridgesResponse": nil,

		// clob
		"/nemo_network.clob.MsgProposedOperations": &clobtypes.MsgProposedOperations{
			OperationsQueue: make([]clobtypes.OperationRaw, 0),
		},
		"/nemo_network.clob.MsgProposedOperationsResponse": nil,

		// perpetuals
		"/nemo_network.perpetuals.MsgAddPremiumVotes": &perptypes.MsgAddPremiumVotes{
			Votes: []perptypes.FundingPremium{
				{PerpetualId: 0, PremiumPpm: 1_000},
			},
		},
		"/nemo_network.perpetuals.MsgAddPremiumVotesResponse": nil,

		// prices
		"/nemo_network.prices.MsgUpdateMarketPrices": &pricestypes.MsgUpdateMarketPrices{
			MarketPriceUpdates: []*pricestypes.MsgUpdateMarketPrices_MarketPrice{
				pricestypes.NewMarketPriceUpdate(constants.MarketId0, 123_000),
			},
		},
		"/nemo_network.prices.MsgUpdateMarketPricesResponse": nil,
	}
)
