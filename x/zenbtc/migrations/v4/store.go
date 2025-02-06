package v4

import (
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func ChangePendingMintTxChainIdtoCaip2Id(ctx sdk.Context, pendingMintTxCol collections.Item[types.PendingMintTransactions], createPendingMintTx func(sdk.Context, *types.PendingMintTransaction) error) error {
	p, err := pendingMintTxCol.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil
		}
		return err
	}

	for _, pendingMintTx := range p.Txs {
		switch pendingMintTx.ChainType {
		case types.WalletType_WALLET_TYPE_EVM:
			pendingMintTx.Caip2ChainId = fmt.Sprintf("eip155:%d", pendingMintTx.ChainId)
		}

		if err := createPendingMintTx(ctx, pendingMintTx); err != nil {
			return err
		}
	}

	return nil
}
