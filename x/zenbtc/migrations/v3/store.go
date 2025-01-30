package v3

import (
	"errors"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func RemoveBadTestnetState(ctx sdk.Context, pendingMintTxCol collections.Item[types.PendingMintTransactions], codec codec.BinaryCodec) error {
	p, err := pendingMintTxCol.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil
		}
		return err
	}

	pTxs := types.PendingMintTransactions{}

	for _, pendingMintTx := range p.Txs {
		if pendingMintTx.ChainId != 17000 {
			continue
		}

		pTxs.Txs = append(pTxs.Txs, pendingMintTx)
	}

	if err := pendingMintTxCol.Set(ctx, pTxs); err != nil {
		return err
	}

	return nil
}
