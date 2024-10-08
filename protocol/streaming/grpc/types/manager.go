package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	clobtypes "github.com/nemo-network/v4-chain/protocol/x/clob/types"
)

type GrpcStreamingManager interface {
	Enabled() bool
	Stop()
	// L3+ Orderbook updates.
	Subscribe(
		req clobtypes.StreamOrderbookUpdatesRequest,
		srv clobtypes.Query_StreamOrderbookUpdatesServer,
	) (
		err error,
	)
	InitializeNewGrpcStreams(
		getOrderbookSnapshot func(clobPairId clobtypes.ClobPairId) *clobtypes.OffchainUpdates,
		blockHeight uint32,
		execMode sdk.ExecMode,
	)
	SendSnapshot(
		offchainUpdates *clobtypes.OffchainUpdates,
		subscriptionId uint32,
		blockHeight uint32,
		execMode sdk.ExecMode,
	)
	SendOrderbookUpdates(
		offchainUpdates *clobtypes.OffchainUpdates,
		blockHeight uint32,
		execMode sdk.ExecMode,
	)
	SendOrderbookFillUpdates(
		ctx sdk.Context,
		orderbookFills []clobtypes.StreamOrderbookFill,
		blockHeight uint32,
		execMode sdk.ExecMode,
	)
}
