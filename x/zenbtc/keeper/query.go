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
	var pendingMintResponses []*types.PendingMintTransaction
	if err := k.PendingMintTransactions.Walk(ctx, nil, func(_ uint64, mint types.PendingMintTransaction) (stop bool, err error) {
		pendingMintResponses = append(pendingMintResponses, &mint)
		return false, nil
	}); err != nil {
		return nil, err
	}
	return &types.QueryPendingMintTransactionsResponse{
		PendingMintTransactions: pendingMintResponses,
	}, nil
}

func (k Keeper) QuerySupply(ctx context.Context, req *types.QuerySupplyRequest) (*types.QuerySupplyResponse, error) {
	supply, err := k.Supply.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &types.QuerySupplyResponse{CustodiedBTC: 0, MintedZenBTC: 0}, nil
		}
		return nil, err
	}
	exchangeRate, err := k.GetExchangeRate(sdk.UnwrapSDKContext(ctx))
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &types.QuerySupplyResponse{CustodiedBTC: 0, MintedZenBTC: 0}, nil
		}
		return nil, err
	}
	return &types.QuerySupplyResponse{CustodiedBTC: supply.CustodiedBTC, MintedZenBTC: supply.MintedZenBTC, ExchangeRate: exchangeRate}, nil
}
