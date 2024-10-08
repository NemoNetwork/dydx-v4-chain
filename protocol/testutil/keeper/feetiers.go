package keeper

import (
	dbm "github.com/cosmos/cosmos-db"
	"github.com/nemo-network/v4-chain/protocol/lib"
	"github.com/nemo-network/v4-chain/protocol/mocks"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	delaymsgtypes "github.com/nemo-network/v4-chain/protocol/x/delaymsg/types"
	"github.com/nemo-network/v4-chain/protocol/x/feetiers/keeper"
	"github.com/nemo-network/v4-chain/protocol/x/feetiers/types"
	statskeeper "github.com/nemo-network/v4-chain/protocol/x/stats/keeper"
	vaultkeeper "github.com/nemo-network/v4-chain/protocol/x/vault/keeper"
)

func createFeeTiersKeeper(
	stateStore storetypes.CommitMultiStore,
	statsKeeper *statskeeper.Keeper,
	vaultKeeper *vaultkeeper.Keeper,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(true)

	authorities := []string{
		delaymsgtypes.ModuleAddress.String(),
		lib.GovModuleAddress.String(),
	}
	k := keeper.NewKeeper(
		cdc,
		statsKeeper,
		storeKey,
		authorities,
	)
	k.SetVaultKeeper(vaultKeeper)

	return k, storeKey
}
