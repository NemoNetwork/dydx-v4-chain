package affiliates

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nemo-network/v4-chain/protocol/x/affiliates/keeper"
	"github.com/nemo-network/v4-chain/protocol/x/affiliates/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{}
}
