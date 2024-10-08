package accountplus

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nemo-network/v4-chain/protocol/x/accountplus/keeper"
	"github.com/nemo-network/v4-chain/protocol/x/accountplus/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	err := k.SetGenesisState(ctx, data)
	if err != nil {
		panic(err)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	accounts, err := k.GetAllAccountStates(ctx)
	if err != nil {
		panic(err)
	}
	return &types.GenesisState{
		Accounts: accounts,
	}
}
