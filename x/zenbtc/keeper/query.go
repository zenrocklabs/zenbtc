package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) QueryPendingMintTransactions(ctx context.Context, req *types.QueryPendingMintTransactionsRequest) (*types.QueryPendingMintTransactionsResponse, error) {
	pendingMints, err := k.PendingMintTransactions.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &types.QueryPendingMintTransactionsResponse{PendingMintTransactions: []*types.PendingMintTransactionResponse{}}, nil
		}
		return nil, err
	}
	pendingMintResponses := make([]*types.PendingMintTransactionResponse, 0, len(pendingMints.Txs))
	for _, mint := range pendingMints.Txs {
		pendingMintResponses = append(pendingMintResponses, &types.PendingMintTransactionResponse{
			ChainId:          mint.ChainId,
			ChainType:        mint.ChainType.String(),
			RecipientAddress: mint.RecipientAddress,
			Amount:           mint.Amount,
			Creator:          mint.Creator,
			KeyId:            mint.KeyId,
		})
	}
	return &types.QueryPendingMintTransactionsResponse{PendingMintTransactions: pendingMintResponses}, nil
}

func (k Keeper) QuerySupply(ctx context.Context, req *types.QueryZenBTCSupplyRequest) (*types.QueryZenBTCSupplyResponse, error) {
	supply, err := k.Supply.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &types.QueryZenBTCSupplyResponse{CustodiedBTC: 0, MintedZenBTC: 0}, nil
		}
		return nil, err
	}
	exchangeRate, err := k.GetZenBTCExchangeRate(sdk.UnwrapSDKContext(ctx))
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &types.QueryZenBTCSupplyResponse{CustodiedBTC: 0, MintedZenBTC: 0}, nil
		}
		return nil, err
	}
	return &types.QueryZenBTCSupplyResponse{CustodiedBTC: supply.CustodiedBTC, MintedZenBTC: supply.MintedZenBTC, ExchangeRate: exchangeRate}, nil
}
