package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nemo-network/v4-chain/protocol/lib"
	"github.com/nemo-network/v4-chain/protocol/x/vault/types"
)

type (
	Keeper struct {
		cdc               codec.BinaryCodec
		storeKey          storetypes.StoreKey
		clobKeeper        types.ClobKeeper
		perpetualsKeeper  types.PerpetualsKeeper
		pricesKeeper      types.PricesKeeper
		sendingKeeper     types.SendingKeeper
		subaccountsKeeper types.SubaccountsKeeper
		authorities       map[string]struct{}
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	clobKeeper types.ClobKeeper,
	perpetualsKeeper types.PerpetualsKeeper,
	pricesKeeper types.PricesKeeper,
	sendingKeeper types.SendingKeeper,
	subaccountsKeeper types.SubaccountsKeeper,
	authorities []string,
) *Keeper {
	return &Keeper{
		cdc:               cdc,
		storeKey:          storeKey,
		clobKeeper:        clobKeeper,
		perpetualsKeeper:  perpetualsKeeper,
		pricesKeeper:      pricesKeeper,
		sendingKeeper:     sendingKeeper,
		subaccountsKeeper: subaccountsKeeper,
		authorities:       lib.UniqueSliceToSet(authorities),
	}
}

func (k Keeper) HasAuthority(authority string) bool {
	_, ok := k.authorities[authority]
	return ok
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(log.ModuleKey, fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) InitializeForGenesis(ctx sdk.Context) {}
