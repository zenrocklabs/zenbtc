package v3

import (
	"fmt"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func ChangePendingMintTxChainIdtoCaip2Id(ctx sdk.Context, pendingMintTxCol collections.Item[types.PendingMintTransactions], codec codec.BinaryCodec) error {
	p, err := pendingMintTxCol.Get(ctx)
	if err != nil {
		return err
	}

	pTxs := types.PendingMintTransactions{}

	for _, pendingMintTx := range p.Txs {
		switch pendingMintTx.ChainType {
		case types.WalletType_WALLET_TYPE_EVM:
			pendingMintTx.Caip2ChainId = fmt.Sprintf("eip155:%d", pendingMintTx.ChainId)
		}

		pTxs.Txs = append(pTxs.Txs, pendingMintTx)
	}

	pendingMintTxCol.Set(ctx, pTxs)

	return nil
}
